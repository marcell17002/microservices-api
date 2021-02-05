"use strict";
const bcrypt = require("bcrypt");

module.exports = {
  up: async (queryInterface, Sequelize) => {
    await queryInterface.bulkInsert(
      "users",
      [
        {
          name: "Marcell Antonius",
          profession: "SQA Kompas",
          role: "general",
          email: "marcellantonius18@gmail.com",
          password: await bcrypt.hash("marcell98", 10),
          created_at: new Date(),
          updated_at: new Date(),
        },
        {
          name: "Bella Octaviolla",
          profession: "Creative Preneuer",
          role: "alumnus",
          email: "octaviollabella@gmail.com",
          password: await bcrypt.hash("bella20", 10),
          created_at: new Date(),
          updated_at: new Date(),
        },
      ],
      {}
    );
  },

  down: async (queryInterface, Sequelize) => {
    await queryInterface.bulkDelete("users", null, {});
  },
};
