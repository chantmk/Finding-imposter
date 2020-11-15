module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "log", fields: ["logID", "placeID", "createdAt", "action", ] },
		{ type: "doctor", fields: ["address", "isDoctor", ] },
  ],
};
