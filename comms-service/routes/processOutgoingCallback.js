function processOutgoingCallback(input, callback) {
    console.log("processOutgoingCallback called with input: ", input.request);
  try {
    callback(null, {});
  } catch (error) {
    callback(error, null);
  }
}

module.exports = {
    processOutgoingCallback,
}