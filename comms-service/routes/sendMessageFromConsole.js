function sendMessageFromConsole(input, callback) {
    console.log("sendMessageFromConsole called with input: ", input.request);
  try {
    callback(null, {});
  } catch (error) {
    callback(error, null);
  }
}

module.exports = {
    sendMessageFromConsole,
}