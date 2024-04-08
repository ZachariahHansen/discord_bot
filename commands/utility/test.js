const { SlashCommandBuilder, ModalBuilder, TextInputBuilder, TextInputStyle, ActionRowBuilder } = require('discord.js');

module.exports = {

    run: async({interation}) => {
        const modal = new ModalBuilder()
            .setCustomId('myCounter')
            .setTitle('Counter');
            
        
        const counterContentInput = new TextInputBuilder()
            .setCustomId('counterContent')
            .setPlaceholder('What will you be counting?')
            .setRequired(true)
            .setLabel('Counter Content')
            .setStyle(TextInputStyle.Short);

        const firstActionRow = new ActionRowBuilder().addComponents(counterContentInput);
    },
}