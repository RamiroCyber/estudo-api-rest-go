package configs

import "github.com/spf13/viper"

var cfg *config // variavel para que ngm fora o pacote possa acessar o pcc

type config struct {
	API APIconfig
	DB  DBconfig
}

type APIconfig struct {
	Port string
}

type DBconfig struct {
	Host, Port, User, Pass, Database string
}

func init() { // funcao chamanda sempre no start das aplicacoes
	viper.SetDefault("api.port", "9000") // viper usado para definir valores padroes de config
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error { //funcao que vai carregar o arquivo de conexao com o banco de dados
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")    //. pq a ideia e sempre que ele esteja do lado do binario
	err := viper.ReadInConfig() // leitura do arquivo

	if err != nil { // se o error for diferente de nulo
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { //se o erro for do tipo configfilenotfounderror
			return err
		}
	}
	cfg = new(config) // new e o mesmo que & ele cria um ponteiro da nossa struct

	cfg.API = APIconfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBconfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	return nil
}

func GetDB() DBconfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
