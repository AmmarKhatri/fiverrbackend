const {
  loadPackageDefinition,
  Server,
  ServerCredentials,
} = require("@grpc/grpc-js");
const { loadSync } = require("@grpc/proto-loader");
const { requestAppointment } = require("./routes/requestAppointment");
const { acceptAppointment } = require("./routes/acceptAppointment");
const { rejectAppointment } = require("./routes/rejectAppointment");
const { processOutgoingCallback } = require("./routes/processOutgoingCallback");
const {
  processIncomingSMSCallback,
} = require("./routes/processIncomingSMSCallback");
const {
  processMultiChannelCallback,
} = require("./routes/processMultiChannelCallback");
const { sendMessageFromConsole } = require("./routes/sendMessageFromConsole");
require("dotenv").config();

function startServer() {
  const packageDefinition = loadSync("./comms/comms.proto");
  const commsProto = loadPackageDefinition(packageDefinition);
  const server = new Server();
  if (!commsProto.comms.Communicator) {
    console.error("Failed to create server");
    return;
  }
  server.addService(commsProto.comms.Communicator.service, {
    ProcessOutgoingCallback: processOutgoingCallback,
    ProcessIncomingSMSCallback: processIncomingSMSCallback,
    ProcessMultiChannelCallback: processMultiChannelCallback,
    RequestAppointment: requestAppointment,
    AcceptAppointment: acceptAppointment,
    RejectAppointment: rejectAppointment,
    SendMessageFromConsole: sendMessageFromConsole,
  });
  server.bindAsync(
    `0.0.0.0:${process.env.PORT}`,
    ServerCredentials.createInsecure(),
    (err, port) => {
      if (err) {
        console.error("Failed to bind server:", err);
        return;
      }
      server.start();
      console.log("Server started on port", port);
    }
  );
}

// Start the server
startServer();
