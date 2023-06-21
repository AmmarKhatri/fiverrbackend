function processIncomingSMSCallback(input, callback) {
    console.log("processIncomingSMSCallback called with input: ", input.request);
  try {
    callback(null, {});
  } catch (error) {
    callback(error, null);
  }
}

module.exports = {
    processIncomingSMSCallback,
}