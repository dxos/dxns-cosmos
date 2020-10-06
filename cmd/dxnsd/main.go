package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	abci "github.com/tendermint/tendermint/abci/types"
	tmamino "github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	ethermintclient "github.com/cosmos/ethermint/client"
	ethermintcrypto "github.com/cosmos/ethermint/crypto"

	"github.com/wirelineio/dxns/app"
)

const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cobra.EnableCommandSorting = false

	cdc := app.MakeCodec(app.ModuleBasics)

	tmamino.RegisterKeyType(ethermintcrypto.PubKeySecp256k1{}, ethermintcrypto.PubKeyAminoName)
	tmamino.RegisterKeyType(ethermintcrypto.PrivKeySecp256k1{}, ethermintcrypto.PrivKeyAminoName)

	keys.CryptoCdc = cdc
	genutil.ModuleCdc = cdc
	genutiltypes.ModuleCdc = cdc
	clientkeys.KeysCdc = cdc

	config := sdk.GetConfig()
	app.SetBech32Prefixes(config)
	app.SetBip44CoinType(config)
	config.Seal()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "dxnsd",
		Short:             "DXNS Chain App Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}
	// CLI commands to initialize the chain
	rootCmd.AddCommand(
		ethermintclient.ValidateChainID(
			genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome),
		),
		genutilcli.CollectGenTxsCmd(ctx, cdc, auth.GenesisAccountIterator{}, app.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(ctx, cdc),
		genutilcli.GenTxCmd(
			ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{}, auth.GenesisAccountIterator{},
			app.DefaultNodeHome, app.DefaultCLIHome,
		),
		genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics),
		// AddGenesisAccountCmd allows users to add accounts to the genesis file
		AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome),
		flags.NewCompletionCmd(rootCmd, true),
	)

	// Tendermint node base commands
	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// Add flags for GQL server.
	rootCmd.PersistentFlags().Bool("gql-server", false, "Start GQL server.")
	rootCmd.PersistentFlags().Bool("gql-playground", false, "Enable GQL playground.")
	rootCmd.PersistentFlags().String("gql-playground-api-base", "", "GQL API base path to use in GQL playground.")
	rootCmd.PersistentFlags().String("gql-port", "9473", "Port to use for the GQL server.")
	rootCmd.PersistentFlags().String("log-file", "", "File to tail for GQL 'getLogs' API.")

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "DXNS", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod, 0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewAppInit(
		logger,
		db,
		traceStore,
		true,
		map[int64]bool{},
		0,
		baseapp.SetPruning(storetypes.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	dxnsChain := app.NewAppInit(logger, db, traceStore, true, map[int64]bool{}, 0)

	if height != -1 {
		err := dxnsChain.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
	}

	return dxnsChain.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
