package config

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/cloudflare/circl/sign/ed448"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Key                 *KeyConfig    `yaml:"key"`
	P2P                 *P2PConfig    `yaml:"p2p"`
	Engine              *EngineConfig `yaml:"engine"`
	DB                  *DBConfig     `yaml:"db"`
	ListenGRPCMultiaddr string        `yaml:"listenGrpcMultiaddr"`
	ListenRestMultiaddr string        `yaml:"listenRESTMultiaddr"`
	LogFile             string        `yaml:"logFile"`
}

func NewConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	d := yaml.NewDecoder(file)
	config := &Config{}

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func LoadConfig(configPath string, proverKey string) (*Config, error) {
	info, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		fmt.Println("Creating config directory " + configPath)
		if err = os.Mkdir(configPath, fs.FileMode(0700)); err != nil {
			panic(err)
		}
	} else {
		if err != nil {
			panic(err)
		}

		if !info.IsDir() {
			panic(configPath + " is not a directory")
		}
	}

	file, err := os.Open(filepath.Join(configPath, "config.yml"))
	saveDefaults := false
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			saveDefaults = true
		} else {
			return nil, err
		}
	}

	bootstrapPeers := []string{
		"/dns/bootstrap.quilibrium.com/udp/8336/quic/p2p/QmUhm9iZVruSxyavjoPLCfuoRG94SGQEkfxEEoukEZmD5B",
		"/ip4/204.186.74.47/udp/8317/quic/p2p/Qmd233pLUDvcDW3ama27usfbG1HxKNh1V9dmWVW1SXp1pd",
		"/ip4/13.237.250.230/udp/8317/quic/p2p/QmUwMFVZbuqr3vjqYUfihN4v82W4aoqpcu84aLNEo1YFsB",
		"/ip4/13.236.219.103/udp/8317/quic/p2p/QmczCDFuJHZLwpKDme7GsPg3sFmkrmKSX21ez8yFSw6vwH",
		"/ip4/204.186.74.46/udp/8316/quic/p2p/QmeqBjm3iX7sdTieyto1gys5ruQrQNPKfaTGcVQQWJPYDV",
		"/ip4/186.233.184.181/udp/8336/quic/p2p/QmW6QDvKuYqJYYMP5tMZSp12X3nexywK28tZNgqtqNpEDL",
		"/dns/quil.zanshindojo.org/udp/8336/quic/p2p/QmXbbmtS5D12rEc4HWiHWr6e83SCE4jeThPP4VJpAQPvXq",
		"/ip4/195.15.213.171/udp/8336/quic/p2p/Qme5PeqPSNRAvgWYbKoW9yy8S9tUwx1RBGaJ4ByE1a4rnu",
		"/ip4/144.76.104.93/udp/8336/quic/p2p/QmZejZ8DBGQ6foX9recW73GA6TqL6hCMX9ETWWW1Fb8xtx",
		"/ip4/86.106.89.131/udp/8336/quic/p2p/QmU6kHqghuZbLGDivFh2TQw73vQzsNtYoegwhKTRw77R5p",
		"/ip4/103.79.113.68/udp/8336/quic/p2p/QmSheQ43HuLhVYxdAByUV5pskihWRFpADnivuPf4cShZKq",
		"/ip4/207.246.81.38/udp/8336/quic/p2p/QmPBYgDy7snHon7PAn8nv1shApQBQz1iHb2sBBS8QSgQwW",
		"/dns/abyssia.fr/udp/8336/quic/p2p/QmS7C1UhN8nvzLJgFFf1uspMRrXjJqThHNN6AyEXp6oVUB",
		"/ip4/51.15.18.247/udp/8336/quic/p2p/QmYVaHXdFmHFeTa6oPixgjMVag6Ex7gLjE559ejJddwqzu",
	}
	genesisSeed := "5768656e20696e2074686520436f75727365206f6620746563686e6f6c6f676963616c206576656e74732c206974206265636f6d6573206e656365737361727920666f72206f6e65206e6574776f726b0a746f20646973736f6c766520746865202070726f746f636f6c2062616e647320776869636820206861766520636f6e6e6563746564207468656d20207769746820616e6f746865722c20616e6420746f0a617373756d6520616d6f6e672074686520706f77657273206f6620746865207765622c202074686520736570617261746520616e6420657175616c2073746174696f6e20746f207768696368207468650a4c617773206f6620546563686e6f6c6f677920616e6420206f6620546563686e6f6c6f6779277320537465776172647320656e7469746c65207468656d2c20206120646563656e7420726573706563740a746f20746865206f70696e696f6e73206f6620416c6c20726571756972657320207468617420746865792073686f756c64206465636c617265207468652063617573657320776869636820696d70656c0a7468656d20746f207468652073657061726174696f6e2e0a0a576520686f6c642074686573652074727574687320746f2062652073656c662d65766964656e742c207468617420616c6c20636f676e697469766520656e7469746965732061726520637265617465640a657175616c2c2074686174202074686579206172652020656e646f7765642062792020746865206f7264657220206f66206c696665202077697468206365727461696e2020756e616c69656e61626c650a5269676874732c2020746861742020616d6f6e6720207468657365202061726520204c6966652c20204c6962657274792c2020507269766163792020616e64202050726f70657274792e2020546861740a686973746f726963616c6c7920746f20207365637572652020746865736520207269676874732c202050726f746f636f6c7320206861766520206265656e2020696e737469747574656420616d6f6e670a656e7469746965732c2020617320476f7665726e6d656e74732068617665206265656e2020696e737469747574656420616d6f6e67204d656e2c2020686176696e6720646572697665642074686569720a6a75737420706f776572732066726f6d2074686520636f6e73656e74206f662074686520676f7665726e65642c2054686174207768656e6576657220616e7920466f726d206f662050726f746f636f6c0a6265636f6d6573206465737472756374697665206f6620746865736520656e64732c20697420697320746865205269676874206f6620416c6c20746f20616c746572206f7220746f2061626f6c6973680a69742c20616e6420746f20696e73746974757465206e65772050726f746f636f6c732c206c6179696e672069747320666f756e646174696f6e206f6e2073756368207072696e6369706c657320616e640a6f7267616e697a696e672069747320706f7765727320696e207375636820666f726d2c2020617320746f207468656d207368616c6c207365656d206d6f7374206c696b656c7920746f206566666563740a74686569722053616665747920616e642048617070696e6573732e2050727564656e63652c2020696e646565642c202077696c6c206469637461746520746861742050726f746f636f6c73206c6f6e670a65737461626c6973686564202073686f756c6420206e6f742020626520206368616e6765642020666f7220206c696768742020616e6420207472616e7369656e7420206361757365733b202020616e640a6163636f7264696e676c7920616c6c20657870657269656e6365206861746820736865776e2c2074686174206d616e6b696e6420617265206d6f726520646973706f73656420746f207375666665722c0a7768696c65206576696c73206172652073756666657261626c652c207468616e20746f20207269676874207468656d73656c7665732062792061626f6c697368696e672074686520666f726d7320746f0a776869636820746865792020617265206163637573746f6d65642e2020427574207768656e202061206c6f6e6720747261696e20206f66206162757365732020616e642075737572706174696f6e732c0a7075727375696e672020696e7661726961626c7920207468652073616d6520204f626a656374206576696e6365732020612064657369676e2020746f2072656475636520207468656d2020756e6465720a6162736f6c7574652020446573706f7469736d2c20206974206973202074686569722072696768742c202069742069732020746865697220647574792c2020746f207468726f77206f666620737563680a50726f746f636f6c732c2020616e6420746f2070726f76696465206e6577204775617264732020666f72207468656972206675747572652073656375726974792e20205375636820686173206265656e0a746865202070617469656e742020737566666572616e636520206f66202074686520207765623b2020616e6420207375636820697320206e6f772074686520206e6563657373697479202077686963680a636f6e73747261696e73207468656d20746f20616c74657220746865697220666f726d65722053797374656d73206f66204e6574776f726b696e672e202054686520686973746f7279206f66207468650a70726573656e7420496e7465726e6574206973206120686973746f727920206f6620726570656174656420696e6a757269657320616e642075737572706174696f6e732c2020616c6c20686176696e670a696e2064697265637420206f626a65637420746865202065737461626c6973686d656e74206f662020616e20206162736f6c7574652020547972616e6e79206f7665722020746865204e6574776f726b0a53746174652e20546f2070726f766520746869732c206c6574204661637473206265207375626d697474656420746f20612063616e64696420776f726c642e0a0a41742074686520506879736963616c206c617965722c20496e7465726e657420536572766963652050726f7669646572732068617665207265667573656420746865697220417373656e7420746f0a46726565205370656563682c20746865206d6f73742077686f6c65736f6d6520616e64206e656365737361727920666f7220746865207075626c696320676f6f643a0a202020202020202068747470733a2f2f7777772e74686576657267652e636f6d2f32333732373233382f6e65742d6e65757472616c6974792d686973746f72792d6663632d6c656769736c6174696f6e0a0a4174207468652044617461204c696e6b206c617965722c20496e667275737472756374757265206173206120536572766963652050726f766964657273206861766520666f7262696464656e0a746865697220557365727320746f20667265656c792061646865726520746f20746865204c6177732c20756e6c6573732073757370656e64656420696e207468656972206f7065726174696f6e0a74696c6c20746865697220417373656e742073686f756c64206265206f627461696e65643b20616e64207768656e20736f2073757370656e6465642c207468657920686176652075747465726c790a6e65676c656374656420746f20617474656e6420746f207468656d2e0a2020202068747470733a2f2f7777772e636e62632e636f6d2f323032312f30312f31362f686f772d7061726c65722d6465706c6174666f726d696e672d73686f77732d706f7765722d6f662d636c6f750a20202020642d70726f7669646572732e68746d6c0a0a417420746865204e6574776f726b206c617965722c20476f7665726e6d656e74732068617665207265667573656420746f2070617373206f74686572204c61777320666f72207468650a6163636f6d6d6f646174696f6e206f66207365637572696e672074686520526967687473206f66207468652050656f706c65206f6e6c696e652c20696e2069747320737465616420616374696e670a636f6e747261727920746f20746865204c617773207375636820746861742074686f73652070656f706c6520776f756c642072656c696e717569736820746865207269676874206f660a526570726573656e746174696f6e20696e2074686520546f706963206f66205072697661637920616e642053656375726974792c206120726967687420696e657374696d61626c6520746f0a7468656d20616e6420666f726d696461626c6520746f20747972616e7473206f6e6c792e0a2020202068747470733a2f2f7777772e6566662e6f72672f646565706c696e6b732f323032332f30392f746f6461792d756b2d7061726c69616d656e742d756e6465726d696e65642d707269766163790a202020202d73656375726974792d616e642d66726565646f6d2d616c6c2d696e7465726e65742d75736572730a0a417420746865205472616e73706f7274206c617965722c20556e6163636f756e7461626c6520456e7472656e6368656420506f7765727320686176652063616c6c656420746f6765746865720a617373656d626c696573206f6620696e74657264696374696f6e20617420706c6163657320756e757375616c2c20756e636f6d666f727461626c652c20616e642064697374616e742066726f6d0a746865206465706f7369746f7279206f66207468656972207075626c6963205265636f726473206f7220696e206120666f726d2061636365737369626c6520746f206163636f756e7461626c650a696e717569736974696f6e2c20666f722074686520736f6c6520707572706f7365206f6620696e7661736976656c79207363727574696e697a696e6720746865206c69766573206f66207468650a50656f706c652e0a202020202020202020202020202020202020202020202020202020202020202020202020202020202068747470733a2f2f656e2e77696b6970656469612e6f72672f77696b692f526f6f6d5f363431410a0a4174207468652053657373696f6e206c617965722c204f7264657273206861766520646973736f6c7665642074686520726967687473206f66207468652070656f706c652072657065617465646c792c0a666f72206f70706f73696e6720616e6420636f6e737472756374696e6720636f756e7465726d65617375726573206f6e207468697320696e64657264696374696f6e2e0a2020202068747470733a2f2f7777772e6566662e6f72672f646565706c696e6b732f323030372f31302f71776573742d63656f2d6e73612d70756e69736865642d71776573742d7265667573696e672d0a2020202070617274696369706174652d696c6c6567616c2d7375727665696c6c616e63652d7072652d392d31310a0a4174207468652050726573656e746174696f6e206c617965722c2074686520537061726b73206f66206e657720436f6e7363696f75736e6573732061726520756e6465722064656261746520746f0a6265207375726d697365642064616e6765726f757320616e6420696e6a7572696f757320746f204d616e2c20616e642066726f6d2074686973206465626174652061726775656420746861742069740a73686f756c64206265207375626a65637420746f2063656e736f727368697020616e64207365636f6e6420636c617373656420746f206120736c6176652073746174757320756e6465720a636f72706f72617465206d6173746572732e0a2020202068747470733a2f2f76656e74757265626561742e636f6d2f61692f776974682d612d776176652d6f662d6e65772d6c6c6d732d6f70656e2d736f757263652d61692d69732d686176696e672d0a20202020612d6d6f6d656e742d616e642d612d7265642d686f742d6465626174652f0a0a496e206576657279207374616765206f6620207468657365204f707072657373696f6e732057652068617665205065746974696f6e65642020666f72205265647265737320696e20746865206d6f73740a68756d626c65207465726d733a204f7572207265706561746564205065746974696f6e732068617665206265656e20616e737765726564206f6e6c7920627920726570656174656420696e6a7572792e0a41205072696e63652077686f7365206368617261637465722069732074687573206d61726b65642020627920657665727920616374207768696368206d617920646566696e65206120547972616e742c0a697320756e66697420746f206265207468652072756c6572206f6620612066726565206e6574776f726b2e0a0a4e6f722068617665205765206265656e2077616e74696e672020696e20617474656e74696f6e7320746f206f7572204c7564646974652020627265746872656e2e2057652068617665207761726e65640a7468656d2066726f6d202074696d6520746f2074696d65206f6620617474656d7074732020627920746865697220766172696f7573206c656769736c6174757265732020746f20657874656e6420616e0a756e77617272616e7461626c65206a7572697364696374696f6e206f7665722075732e202020576520686176652072656d696e646564207468656d206f66207468652063697263756d7374616e6365730a6f66206f75722020646576656c6f706d656e7420616e6420656e6c69676874656e6d656e742020686572652e202057652068617665202061707065616c656420746f20207468656972206e61746976650a6a75737469636520616e64206d61676e616e696d6974792c2020616e6420776520686176652020636f6e6a7572656420207468656d20627920746865207469657320206f66206f757220636f6d6d6f6e0a6b696e647265642020746f2064697361766f77202074686573652075737572706174696f6e732c202077686963682c2020776f756c6420696e6576697461626c792020696e74657272757074206f75720a636f6e6e656374696f6e7320616e6420636f72726573706f6e64656e63652e20205468657920746f6f2068617665206265656e206465616620746f2074686520766f696365206f66206a7573746963650a616e6420206f6620636f6e73616e6775696e6974792e20205765206d7573742c207468657265666f72652c202061637175696573636520696e2020746865206e65636573736974792c202077686963680a64656e6f756e63657320206f7572202053657061726174696f6e2c2020616e642020686f6c6420207468656d2c2020617320776520686f6c6420207468652072657374206f6620206d616e6b696e642c0a446973636f6e6e656374656420696e204f70706f736974696f6e2c20696e2050656163652050656572732e0a0a57652c20207468657265666f72652c202074686520526570726573656e74617469766573206f6620746865206e657720496e7465726e65742c2020696e20746865206669727374207068617365206f660a74686520477265617420436572656d6f6e792c2020417373656d626c65642c202061707065616c696e6720746f207468652053757072656d65204a75646765206f662074686520776f726c6420666f720a74686520726563746974756465206f66206f757220696e74656e74696f6e732c2020646f2c2020696e20746865204e616d652c20616e6420627920417574686f72697479206f662074686520676f6f640a5374657761726473206f6620746865205765622c20736f6c656d6e6c79207075626c69736820616e64206465636c6172652c20205468617420746865205175696c69627269756d2050726f746f636f6c0a69732c20616e64206f66205269676874206f7567687420746f206265204672656520616e6420496e646570656e64656e743b202074686174206974206973204162736f6c7665642066726f6d207468650a64656d616e6473206f6620756e656c65637465642c20756e2d436f6e737469747574696f6e616c20666f72636573206f6e20746865207765622c20616e64207468617420746865206f6e6c79204c61770a77686963682020676f7665726e73202069742069732020746865204c617720207768696368207765202068617665207072657365727665642020666f722074686520206265747465726d656e74206f660a736f6369657479207468726f7567682074686520726570656174656420656e666f7263656d656e74206f662074686520436f6e737469747574696f6e3b20616e642074686174206173206120467265650a616e642020496e646570656e64656e742050726f746f636f6c2c2020697420686173202066756c6c20506f77657220746f2020446973636f6e6e6563742c20206a6f696e20696e2050656572696e672c0a636f6e747261637420436f6d7075746174696f6e2c2020656e6761676520696e20436f6d6d657263652c2020616e6420746f20646f20616c6c206f7468657220204163747320616e64205468696e67730a77686963682020496e646570656e64656e7420204e6574776f726b7320206d6179206f662020726967687420646f2e2020416e6420666f7220207468652020737570706f727420206f662020746869730a4465636c61726174696f6e2c2020776974682061206669726d2072656c69616e6365206f6e20207468652070726f74656374696f6e206f6620436f6e737469747574696f6e616c204c61772c202077650a6d757475616c6c792020706c656467652020746f20206561636820206f7468657220206f757220204d696e64732c206f757220204d616368696e65732020616e6420206f7572207374726f6e676573740a43727970746f6772617068792e"

	config := &Config{
		DB: &DBConfig{
			Path: configPath + "/store",
		},
		Key: &KeyConfig{
			KeyStore: KeyManagerTypeFile,
			KeyStoreFile: &KeyStoreFileConfig{
				Path: filepath.Join(configPath, "keys.yml"),
			},
		},
		P2P: &P2PConfig{
			ListenMultiaddr: "/ip4/0.0.0.0/udp/8336/quic",
			BootstrapPeers:  bootstrapPeers,
			PeerPrivKey:     "",
		},
		Engine: &EngineConfig{
			ProvingKeyId:         "default-proving-key",
			Filter:               "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff",
			GenesisSeed:          genesisSeed,
			MaxFrames:            -1,
			PendingCommitWorkers: 4,
		},
	}

	if saveDefaults {
		fmt.Println("Generating default config...")
		fmt.Println("Generating random host key...")
		privkey, _, err := crypto.GenerateEd448Key(rand.Reader)
		if err != nil {
			panic(err)
		}

		hostKey, err := privkey.Raw()
		if err != nil {
			panic(err)
		}

		config.P2P.PeerPrivKey = hex.EncodeToString(hostKey)

		fmt.Println("Generating keystore key...")
		keystoreKey := make([]byte, 32)
		if _, err := rand.Read(keystoreKey); err != nil {
			panic(err)
		}

		config.Key.KeyStoreFile.EncryptionKey = hex.EncodeToString(keystoreKey)

		fmt.Println("Saving config...")
		if err = SaveConfig(configPath, config); err != nil {
			panic(err)
		}

		keyfile, err := os.OpenFile(
			filepath.Join(configPath, "keys.yml"),
			os.O_CREATE|os.O_RDWR,
			fs.FileMode(0700),
		)
		if err != nil {
			panic(err)
		}

		if proverKey != "" {
			provingKey, err := hex.DecodeString(proverKey)
			if err != nil {
				panic(err)
			}

			iv := [12]byte{}
			rand.Read(iv[:])
			aesCipher, err := aes.NewCipher(keystoreKey)
			if err != nil {
				return nil, errors.Wrap(err, "could not construct cipher")
			}

			gcm, err := cipher.NewGCM(aesCipher)
			if err != nil {
				return nil, errors.Wrap(err, "could not construct block")
			}

			ciphertext := gcm.Seal(nil, iv[:], provingKey, nil)
			ciphertext = append(append([]byte{}, iv[:]...), ciphertext...)

			provingPubKey := ed448.PrivateKey(provingKey).Public().(ed448.PublicKey)

			keyfile.Write([]byte(
				"default-proving-key:\n  id: default-proving-key\n" +
					"  type: 0\n  privateKey: " + hex.EncodeToString(ciphertext) + "\n" +
					"  publicKey: " + hex.EncodeToString(provingPubKey) + "\n"))
		} else {
			keyfile.Write([]byte("null:\n"))
		}

		keyfile.Close()

		if file, err = os.Open(
			filepath.Join(configPath, "config.yml"),
		); err != nil {
			panic(err)
		}
	}

	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	if config.Engine.GenesisSeed == "00" {
		config.Engine.GenesisSeed = genesisSeed
	}

	if len(config.P2P.BootstrapPeers) == 0 {
		config.P2P.BootstrapPeers = bootstrapPeers
	}

	return config, nil
}

func SaveConfig(configPath string, config *Config) error {
	file, err := os.OpenFile(
		filepath.Join(configPath, "config.yml"),
		os.O_CREATE|os.O_RDWR,
		os.FileMode(0600),
	)
	if err != nil {
		return err
	}

	defer file.Close()

	d := yaml.NewEncoder(file)

	if err := d.Encode(config); err != nil {
		return err
	}

	return nil
}
