function rejectAppointment(input, callback) {
    console.log("rejectAppointment called with input: ", input.request);
  try {
    callback(null, {});
  } catch (error) {
    callback(error, null);
  }
}

module.exports = {
    rejectAppointment,
}