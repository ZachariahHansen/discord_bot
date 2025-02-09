const { Events } = require('discord.js');

module.exports = {
	name: Events.InteractionCreate,
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
