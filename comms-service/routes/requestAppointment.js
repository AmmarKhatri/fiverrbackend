function requestAppointment(input, callback) {
    console.log("requestAppointment called with input: ", input.request);
  try {
    callback(null, {
    //   provider_id: 0,
    //   customer_id: 0,
    //   recipient_ids: [],
    //   appointment_type_id: 0,
    //   startDate: "",
    //   endDate: "",
    });
  } catch (error) {
    callback(error, null);
  }
}

module.exports = {
    requestAppointment,
}