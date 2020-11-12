module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "quarantine", fields: ["user_id", "start_at", "end_at", ] },
		{ type: "covid", fields: ["status", "created_at", "pub_key", ] },
		{ type: "log", fields: ["place_id", "check_in_at", "check_out_at", ] },
  ],
};
