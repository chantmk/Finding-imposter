module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "quarantine", fields: ["id", "user_id", "start_at", "end_at", ] },
		{ type: "patient", fields: ["id", "user_id", "status", ] },
		{ type: "Log", fields: ["id", "place_id", "user_id", "checkin_at", "checkout_at", ] },
  ],
};
