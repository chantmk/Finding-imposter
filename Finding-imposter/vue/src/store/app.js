module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "quarantine", fields: ["startAt", "endAt", ] },
		{ type: "covid", fields: ["covidID", "createdAt", "status", ] },
		{ type: "log", fields: ["logID", "placeID", "createdAt", "action", ] },
		{ type: "doctor", fields: ["address", "isDoctor", ] },
  ],
};
