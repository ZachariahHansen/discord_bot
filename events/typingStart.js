const { Events } = require('discord.js');

module.exports = {
    name: Events.TypingStart,
    async execute(client) {
        console.log(`test`);
    },
};
