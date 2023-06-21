function processMultiChannelCallback(input, callback) {
    console.log("processMultiChannelCallback called with input: ", input.request);
  try {
    callback(null, {});
  } catch (error) {
    callback(error, null);
  }
}

module.exports = {
    processMultiChannelCallback,
}