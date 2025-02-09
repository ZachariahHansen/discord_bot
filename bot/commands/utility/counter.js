const { SlashCommandBuilder, ModalBuilder, TextInputBuilder, TextInputStyle, ActionRowBuilder } = require('discord.js');

// https://discordjs.guide/interactions/modals.html#building-and-responding-with-modals

module.exports = {
	data: new SlashCommandBuilder()
		.setName('counter')
		.setDescription('Creates a new counter.'),

        async execute(interaction) {
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

            modal.addComponents(firstActionRow);

            await interaction.showModal(modal);
        },
};
