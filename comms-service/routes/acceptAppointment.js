function acceptAppointment(input, callback) {
    console.log("acceptAppointment called with input: ", input.request);
  try {
    callback(null, {});
  } catch (error) {
    callback(error, null);
  }
}

module.exports = {
    acceptAppointment,
}