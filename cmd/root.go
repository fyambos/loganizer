package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd représente la commande de base lorsqu'elle est appelée sans sous-commande.
// C'est une variable globale (dans le package cmd) qui est un pointeur vers une instance de cobra.Command.

var rootCmd = &cobra.Command{
	Use:   "loganizer",                                                                                       // Définit comment la commande principale doit être appelée (ex: 'gowatcher').
	Short: "Gowatcher est un outil pour vérifier l'accessibilité des URLs.",                                  // Une courte description de la commande, affichée dans l'aide.
	Long:  `Un outil CLI en Go pour vérifier l'état d'URLs, gérer la concurrence et exporter les résultats.`, // Une description plus longue, affichée avec 'gowatcher help'.
	// Run: func(cmd *cobra.Command, args []string) {}, // Pas de Run direct pour la root si elle a des sous-commandes

	// Execute ajoute toutes les commandes enfants à la commande racine et définit également les drapeaux.
	// C'est appelé par main.main(). Il n'y a qu'une seule fois Execute.
}
