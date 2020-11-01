module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "patient", fields: ["status", "user_id", ] },
		{ type: "log", fields: ["place_id", "check_in_at", "check_out_at", ] },
  ],
};
