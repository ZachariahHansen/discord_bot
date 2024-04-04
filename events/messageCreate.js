const { Events } = require('discord.js');

module.exports = {

    name: Events.MessageCreate,
    async execute(message, client) {
        if (message.author.bot) return;
        if ((message.content).toLowerCase().includes("cum")){
            message.react("<a:CUMDETECTED:1225282939147911209>");
        }

        
    },
    
};