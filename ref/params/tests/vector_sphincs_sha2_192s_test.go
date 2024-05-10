//go:build cgo && ((linux && amd64) || (darwin && amd64) || (darwin && arm64) || (windows && amd64)) && sphincs_sha2_192s

package tests

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
  sphincsplus "github.com/ioerror/sphincsplus/ref"
)

func TestSignVerifyVectorSpecifics(t *testing.T) {

	privKey := new(sphincsplus.PrivateKey)
	privKeyBytes, err := hex.DecodeString("7894d3acddbdc67eaf2320994ff21a9a38bccc1774955ed540e57c54431917deb11a8cd9df87856a8edaa96a27d4f044e6ffa87a964148a4ba38a19d15ce8d0d2a3b964e5bb6dd1552390abb350079d44aebee256aca81859b69063f9e909563")
	require.NoError(t, err)
	err = privKey.FromBytes(privKeyBytes)
	require.NoError(t, err)

	pubKey := new(sphincsplus.PublicKey)
	pubKeyBytes, err := hex.DecodeString("e6ffa87a964148a4ba38a19d15ce8d0d2a3b964e5bb6dd1552390abb350079d44aebee256aca81859b69063f9e909563")
	require.NoError(t, err)
	err = pubKey.FromBytes(pubKeyBytes)

	signature, err := hex.DecodeString("695e2abd42d4f57a5d6ae33e9db433fa95fb3820d397d1bc8512dcac09ba07a4a041d18593cb20b4553c50ac7458b0126f9300f20d66907478c1303d165a761b53469815f62f35cd87e88c07a47a9bbc141fdc8e76ecd82c80ca59dba5ef95a6e51995ed8d442f00de86daa27805876b64b192ceb912166c6ce9b6806b935e292095f63cc252fd293399ed17b1e2e128f07f8dd7c2c114ac9a0f4f1537bd72918b1920363057608aeec570173b364702fcd7035699b244150134aa10fb9ca750f12b7db6be172e584b4ca6bf8d56de10fdc3b14bd99bdd9c67ec5466e2c8ddd1ae6c34baabb2e221e053c13455f97707b1edd0572c0d173ec5cccab1531441ff6ed36cd7184500d878f579ee47d2845c1930c0b395fe968af35fd7b98f64b7e2e27384da9b92e07e7af6398e611c4374d57a9e068d429dd0abd81cfd1d394d672bc010d413d0925de0986c7fd933b91b85613ddf0b7d1935da82a07c6d4c3e2d61e47d38eb089e356db817d71f2f6fbeb078734ab937bc7785693efefcdb8fea336b23f7c69b490836184494eb98cbb346524c9185da14b8b8a0e6fe0c87f779998423d8e912023c822f2fb8084d9507264aec9b3f12d72e784cfaab2043953482117bc02ceee8e57bc79d1f1748edbec7d376fb6bc0e7a7b6b4eb87c3d82d6116c40e8651ffad3715cf7cbfef2f2a8ee0e69f330647dd7d38eb918beaae100abb22682be5d61f21b5e484e85bb25bbd02abea7bcb6cdb4a535455998b5ca70b644d35568313cdbdc95f418ceed73c01628eb69d1ffb4c9dc37d0e5d0cefa8ce1d3e1fc23e44a805d82c6016c47428c9f3bf141a22acab1095697a376a70704004a37f164a707412e04d3dd79e9318fcc43ac16540085d897df36f20f6552b609297e1cb4b2bce6ecf00b67ca3ef836b8681792b0fa59a274a802ef05d8e32a92e418c65634706761c0094eac4a77eb80323f5d6d71d01053fdf3e1313f6bc15cb46c6738fdc3baab6ae73a06ccb53956e470eb67b8c64522f124d4c2361830f5e7c797f78c2ab5c2cce27419b474611113a96d3b21835482934fbaa4d89b466b2d5479e012d6017296decc8edd408e1ecd0d319fc154aa726f70a01bee3bce5ff0f6d6d4dac52777567a298c92ca7baa91aef9e72511cbd9beee9ac1c224e85d302f4934b4c090a58132b4792eb103e3ca02da73c5231e088ca6e4af3e2f224735d8c69a9e4fc3db7fecf475a3fd228f95dd2fb8cf7d8a07e2b5d024653ec7b59e3950b31eac8a2f8ef66be19eaf5245d1e019dbebf30051b16cd5c061a551f68d834fb807759481af95009735b9314c0ec1874134abe11fe373677c22bd569780f893d10071b3fbfa3280b9b96a550efb7598c80dd6507d405934d78aae0dde261aee300d3e43679c7abc7d5a68abd4122bd9b84fb68de63050452bf5bc2519219745e7b275ff5d90f3547cabfd4e472155d3b58f0dc578a39d1629b856d09c502e3b435d352d13810c23cd1fc5dbf78aec44c36a818eab5a8e07ab135c30569c1f99d6a12c48cb37d36ef1276502d9b51b57cf0d3b1d0661716a714cd0ad167ad21916da27fa5327362ee543261bc5101936b7997ccc74fed544beb709d180f59a026acc649f9bd8a7593c1a9ee3a0c9b23d3dbc2405a7209e20c94c6affe29051cf40cf4ed6240f848dc59c29eb8429760c889b3770c3d4d5f248270d7a7fe9e1c16c240e43eea4b7869dabd54bd1b2898473eedefe092a5446b4151327f4b3121f20b5beb47c60b8f5ce5abbcdfb5444b1e2bfbec2e96db2dcb7f0f305aa57756db89f44fb8ef6385769765783b89b96708416dd9700168b2d3cbc5216aa41d0e0444294a6c44758cc7f05ea813eac0dc5845ae6353645a89d261b094e1c519531d0e9736786d8876d0e8a3fc06aacba94e55a1c513ad721ec62a6165fd7ff704814c71728a83d783ea1879ef3b8e6265591481db74d7e0394043874a204a55b7428cc1d867b28706d98144ca0c047e8b4b289edad869470415586957cb38698deafac80942c0f1f74454f871d19480a2837c11130daa22e26a587a2b59a45e1e1cdd2a90283bcad39417c2c16efa841295b4c368c326f6b7b0d6b32323036fd598ce0aca02765b289d80c37cb76ecc85f703c5479e3ade949de3ed554468612bdbbedf4e8301037afe2efced2afcb7ff1dd3b0bb0d15faadf7d4c9adff58e8f5fa000041d91ab7874c0b1b2b6e3e300bd316e08daa51a115740e3056907b2d8de176e9ac26289e58aca49ac2139388c4632dccaa50ed53ca3b92b7591e72b5059ce546a76dcf7b79a3de9b96576f63a63774f2b20f0a6e3dca5c46072322daf691d18927b5218112c8e1cc656449ff3ca87611086a7b10c7681c6d9c22c403cfd2f0518d324dfe27ca752183b379a4de566a42a265c3c4423dd4c4c2b7628ce7cfb9bea2d512254114d98d3646e72ebd93990dd1d58e25b2fae43be0de767327136789ca0118947164b896d245ec0c306bb05d1959c51eb2f2769201d584271ea50c0adf14142573821a26606479019f69597c2e8198a184bef99f2c4a605418f01939e75a19c3c15959878f88b85aab957e668a155b9e3110ea786937a2dd8f81c103edda9450f08ad65fd16ab7f7a56185d768c1175b88c5adaf3152f12dd8393bc6f7b50ced365480d9788b2229fe5e214cfd785d25a6eba3ffb46b36d6e702cc9480365b3072d11409e533e89a99c943d7ba798ea3283c578467a2062aca12713e9ee351af22241cc32645c0e784688fd70bda6478dcc229fe210edb596caefa8d8648b6957622787d00dbf29a62f3b1254cbecf64fe9517d3a30507b313ff27278c416d83a825472d5397b1970372b7ab90b82751b98accd13df948f0e9c68223dc791a98c31b55bd2755234ccff4fc8fdffd58a211a52cd98235111be7da9702b1e4609ac019c80a7381da47a041fe1abd8c84c56f99409246ed7c26f9da56959878cc1bb8e641cb5f6dba11749c1a082add454464cad329682848e6798c07383133d4b0523367ea48a5b75ebb6a3572c0b08ade84b945444601864882fc00a190430291b2339cd732733a8a15d2b5045e2bf445dbdb37aa32edda1c0338c6ad6c0b0f72d410817f88e8172d2f184f107da2888e060a41dff9d8ced2d9ee526af88c54e9385626b83673edd7caebf359402c4e86bd0556528023501fbe8b3b34f4579280c87a0c685a151d2810ca1e03035557a14097f182f9150e16cd228a0579b39454555ae4dcf0b0d4bdb5dd367fa2d347a4673a06be4e4d49f7b130cde3fcaf8cfabd785a5c0f16810fd26cab4793e25bb6b96b823bb02bec18ebef0e1feeb2fca206b1fa8fcbdca6f75e99517c3dfa1b69c85ac6d0e5d41c68bd923ecbe107cd3d09b4ce688972eaab92fba0a98e98b86073aac66a56c9f770b54a8def223ea7fbd7135f190efe2dbaa5b2603e9bebdb3fc956f57dd545f618c4be60526aefc2495935a48f094986ea9f7956780c1b61d5230bdb2acdbe49556944bbfce25a8c0a951be43791297119ddb3b69ab5348859d43a282ce27b34fcbb6cabb70e1afb6d4dfa2d46bc909fb34c5ac5dfcb7c1664f63f1ce7202ae1d0907157ad4b774994e843513d8f044330790707d700baf3b04a9f324097236ed82bcbf5e9461215795bec50bf3ab6d751af279c837ca33b1719ccf0910bef0e4eeac039e7449674008b320a531a74d28e3dfb7829b757e62caedfadb81ca37ab36575b3e5a417c41aba806c658f2d4a877700d1a70e3f88d2778a5a76859d9eb6e2f2ff22e033b2b8b7fd666ece2afd5fad3b6e52b176ca6d81c4602832a9c29ccb518fe8efeaeac6989662e67c75bbd8323a384a1e1f76921c7bdd12d6dfe8778f509c3ff9b2b9fa1e51dd70ccd132e5723227acc7a9625021770abd482246a8662ea23c1c911087fa398f5e33dd50c45312f5cf1867d25bb7f9129e5cee6bd1b25a95dce6a166868cf642ccb0e30ee3fcb025c7e70dbe01bbc07f1c33705426c7e3affca53fff7c2ba66d3988fea017daddc1879a2faf21c4148089c0888782b0f42a3f5b76f3939626ab94e5d3b710685e464e02d9a5113e3b254a28516e252eed0d9e0cf28ba06ba8c00ecbc05d9dc6a9ee3ffac966488f9ad5eb95d7b0fd5e36c89de40f7b62ec3bd98e1d517da12f5c876766d2e8ae54afcb19e1f1d57f334c76b4e7b62e562ac3deb3e5c811ffb8dbf7cb5ac8c0adb2da9be431a0a0776b4a0a9e98530a2fecc606c89bab2b04b9e71664d0bdaf0569aab80a84df931d574983dc3ecc2ada552e8c65137a0c85d5793c542193e2d2e0c649535593fa8e86bd05f3737e4da1c40d466d9e349bb5b13cf17e86abecf8c2e522d4ccd3b409288243db337042beee17cf714b4e29136d7dea8576ebb29c72af7f28d2290f04b921a29ee38a0845b351e1caf5b80020188de43fbc8eb062f42593dc44cfb26d6111e67e65e8cd2440c89e8b75087daf457e1ece97100de4810350dbba57c7014add81f94325b47f7367a0ce533683bb97cf689ebb20920fc4265317feb24d467f9193adc4856987e0481d365c8e03a3d6fa42b2c37b0c8d8d60b394d59a6a6c86768f6885708afaa18b82b045a464c95fe4b8b9221a4974aa9c1a54cd68cf0e9953dd742936ddbffe78482c924a31128a5daaa02ce0d8ef26827ba1b2758437f73a6f366e4d9a25458047a94e9ea5a1274831e2ddf24fc51a907fac8e7356c52ac5da90df88ca42715837e747faabee14e48b2508d54da7ec5f2ee16e6ebbcf7dcbad13cfbdb76024d7bc1c5daea8afe5ce88023d5bfee52a2dd76d77e2dd5f7ebab4b158fdc227efcf1707b715ae7200b1e8125e0caa64ef7bc824ef74c44fb1e071f4a72020472a17b9c7babe9944bc19b7932533cc7bd143babef0bb42be5d34ed7a8748e5671153534b48caccd89ccff5f97be1919a9e2cb1daa1b873d9458d74e9d5b7d87753a0d3881e0e338e25cbf380a21862183b188e0c8ee67e1eae297d780f8a9e9882206add88fb6a56f21699b01837874410fc60368ad3b23026a4856e081fa5b554ec7b4abb3e0536c1fac0aa67603d1b5219fec60c3d4c9714102de568f76315d78b74ad5a39e71682337fcc3ba0ffab45c1255f6a8fffc39012ac56d951a09ceffec9c51278aa40202609c0008c6da41ae3570e800022f9188e4f4a94253b1d2f750d40e8bacc377af305cc17cdb330202875dd35d92fa8c714706f1821b14c4be98635181831787067f3199e480b9b087075f25dbac98548d461073d31b71460a2f6978b3435bef5d359f48f201931b653fd5b4d10b136cc76470a302cb711488a243ec0bdc3dd11a2ae26750cbd3b52f8b9086fd5ae110b1209d09e149028194a9a7a4d7fb26050527c495c3fbee9f44314d75232271b40bbea25650465f24ace75ac17a15b6f0772f88d39937bf77c23c65edbdbe86d7bd64f74f4fe0ffaee2841afaf0dcfa55d2e3abdf947092a164b31b8f2eb90a069fd457fd99c53bfd135fd6b1e359be02c9b1f382d4ac3d343e6386c446839cd5e110b09da4b459f83e4f95e97a950761be9254174f5af748318f84fd9b7b313d8ee0ab3faf817c26b200332d5733b9553e3b498c05cbaa524fbb6604796873ca0e2ce40b0747dc6a445f38078d47705da53986b67320a7015a8b5d3f6853b9f1bd5db89af81305874446d3682f50c42a2a9d643e39c9bec311a1b949934294b52d9533e96293b4be67920116b6c581ceb44a10469ca910c96c9f79c1984a53c6f365e46c743ef1ff31ded06c1d818d94f930cfe1d4810623a62454c9988a427efba4cecf556961f6db1b50a5260dc96055dbee9e6bad55b9ffa67d43d819fd471d0412c850df5d0a8db32e16b6089375595705f955dbc9597ad52ed492168b97cc811bde2eff5721e438c7ce15d45a6daed137cdaafb27b4171619204d8a9a85e577bddabdfb6b4e4506b84b4748df0c3e2e3fe2cb68b748be4d08e7c0eedc961c00f2ca15c004f927c23771b50ea2fb846e41a68ee33be946dad2f0c9744eba85a894793d95270d8406035df48001db757a52ecc1f3cf643d5ae37d0cb34aa2314008ff150f607789722697d9f22bffd6fd5c65b4250358065185652edc7e2198228b72d871b4d1820affdf0219ca60b739ad50fb37d3da0ae86c7ff6425955180647b64114fb7a9aeb73b5b51b7f0e0d110d61cb0af3c8d9d996d6af4081d9075244833ec3e76f19b91f29ae7b5e95dde3e464356204d33fd5815db3d01ca825b612e79c81a4bd530f16c4146bd774b137c67b5168904668bfcd6e54c3b4bb65f0989010983de544b3c73df0cdef17fdd5d6c1a7abc653234915394243021f96687c196c26a2d5acc7e683d76c685fc83f7dc2fb67e6d2208a290472dee1d03dc800bdb311fcee4e3b5cb9fa071af7f8f8170035e2b3ba773b8fa238f82a6653ed9e511557c228d7bfe667716f4c149f754bfae36e4bbac17e357d4ff0c86f4d4a89979d89e171fded88b95dccb24ff524c6aa88ecc082e587c612c457d0110c803df353b51f4ba6efe7c37041a20391cde8d4e60ac3cc96c4efd083c749baf38dd386261bea4deb3afd6b634fee480e831c1f8e7f25286dd139d90bde26438f33aff0745a7a8819ae1ecadefda7a8861f53bf7b54fab95ecaf16700cda0ad5b45dc5860fc50fc6b62ad898ed343ce0b7bbe341fd9d66a7a9b674483a0287bec26f8e24e0f90a2146978a2202f3556a966ddbbe281a55f1ff7e9a384c2bae75934eead2f4bfe0c5e27a118ce2e0c391f15486b4eba6c76a47c93d383e6358b6dd3289a7d37926c59c58a520edbda60b66c46a38657986a8aa057a78b488eb633f50083a1e9e883f6a0a0312fc81cdab82fe08162a537cd317edca87f181844bf51d1d936115d7e80e97c670eb3da21babd6f3bcc1c9e5ae6855729dfff2b3d6090075b741416903a722c1cdc1ee635f6a6e1c53a1c17f54e38ede69fc1b0b663790628862b461835084bd3a47253cb5739e1d77d874a67b53fca0e1866f06adb69ac70718b1118521e7c2d5337dffc7f6c448e01e56c85a4623b67a120840c2d98d0d117cf814ea88fde64aaf0ef5d887b1ba875ffd1ad14c46dbb04116a0a0e12c67ea2d89f61df867f2ace4e93c552903ecee5492891eed8a4583e168a6a2818f4774bee77e8d80059a5c6a9417b6b08b80e0499ee9e915bfb76bc799dbe65e9ab75f45e29adf041f5c12322f319f118d564d008f669a52ca331884a4561e0d37809e168c803db820bbb8f074455210796eee1d9eef745563a84a85a82ea4bb2b140a8285387f6086e04ef7ec66bb84eeb012c1430080c5c0917abf69880351a78540ad57e5859e8984811985bef1bd91f82d84992ecfcf7ef6bd0e709f992fce5cd7e58d789e683b2a1a0d1828b55c03481959d38a35e7c5f2a643c348a8b84df6cf9b8d09db3783a9d46a49cbc442a915a227cbb0afee9b110baf1fa9ba1f46a24cd73f5221268fef55723d8a52a950e2799701f73f9cdaf54a42ccf1a5d14b4be4eba531d273561515a56f0b18eda56d97ae114e48bf2a1ba8ab90bf84303a0f40a083f68e05c9cdf485ead8f7f631ca821f336128fb3891eadb020cbda572c2763c13b392b2e223e25b7e6886a2f84e0d217028875fd275df41003f8c337a532cdd5a32335f8e17dc68d66c225ebeadd128a730e2fc8305f1ce1f2d833db15563abcaf84f6483cea1e7fd8081e59b16d2895cad8deb79d02eca2d5cd37ffb063e391af24cec3a173664dd378bee11a4615564abc17f94af1c52e1fb32aea3fe079c4ed78916ded0dba651b6892cfccc4b1db9cc9d85e0a1006a0fb0ea7e65572e365640af43b8b0409d0077498ccfb8d53c9de8f052af7ddba8a7fefd4707656cd4784245103d84e0c267d5198dbd890e0813e7a63f5a44eb7d8e6da774daea7d7664ed0c35deb27fe2f8802493586862bfc5cf5a0dc85e9b9b9056281949cf092ebc53e17c10f82ba1a29508036145beea977b2a641857418944bf12eea2c5d192995ea444d751819307a91aa81760e41fa52d639005a89c7201ce29009940db3769540fc19882642a207aebbd7199b511a93ad677a3ca16451aed91a32d6ad64aae785b8577d9ae1ffd97f1bff09f534f591d6efe099ce64eb085fc11c3e5ef597ca78c9e174c6f3715416443c7921eb29772e60be0180b13724c383a6dc285336004f9dcf30350cb87b7cd0df335bc831f5cca35ba6d565a1a616fe367ee8ed0bc45500f55355b05367e7d5eef509fd416949969dfd75a3904c93235d690a7e01ada19a7fc5f0dc09b5e67d696b0060d22c075c1c001e394bc7b2452f6ef4e1f336e414c7a4c4f889b5b8594b9a7394e8cd80619b2d53a13b620da4342fe98cf6bc139583821c8bd9684e3a3d9a64ed3528de901bdfd14bbfd72723ee1222a64ab6ee88e0a380b25bd8eb65dd885d463ee6f0fa6870795a761ec21009733c83f1bf4c3bc1b9402577b591dd7779bd41dd817d4b893139861cb216b09d03cb23181813c7fd3fcbe071de91a0e652e9e3c6958722e861b69cebe5f253aa02279d223bdf20e14f9a86ee1d7549a4688c340e7ffbbce629d17b7782d8f3bf8a6d5430102a9c9761eaa3951eaa1d178c97c2fd403199cbacf3c3ad294ab5f65aaf561c32cc20974a71138d5c84f677714925219f39d081166f8b375f9821692e92c0da9ee48729c6c8ed66ea88c081b7bdd40d11ef15adc49ec28c63009ead01c1d7f48cb726d91ed16f91db7540df9f8f2cf9151c7d6a614419bee76700f546681a754b3a64a497e4c26fd42148a4548c03a219e8143bf33eecaf16bd6a75f442b37dc1609e2032731e51ff45fec0e49ffdc5be7f3494cbf1c9bc7e1d92f6a1b9a3c83d9c10579851711e330b0cfa439ce28cda73fe3d709a880291cbfc3b2a8e4173cf8674ac92889c26039b332c6c530448e48911971505a7aeb328360bd2d8a6343e16caf677c095a15792a0e8146ff12934bd143c36d501fe52e1f18c3133df512e425d90428ef6daa4f847217f1b3c5e07e8314ed5e043ac0cbae3dfdcc7d97769293168500158855c61e86ff6e90c942794e2100a29a90cec89b58760ec8c4467d952eb53c9f3e63ebb89017fe2bb69627c0b8fb28032979d5428b1115ab455707521b4ee74c0c30d0a8942c0f7461b3cb17dcb2ded7e8a94360b22d2c12d3b0f675d5bb4b02c78386b000fb9d7e5d691b44e50a9e4d90a58ca6b134fc23d1fa83b432f10300f7f00370146d4a94a1ccc34d5ccd8ef85a5ee63c313890d6c39a35a75e6e595cddabe19cecdc9bd4662b32f3547ae1589ed9af5deb8ba43cb58a7408e47104ada8432c94a16d63ad63b542a947b654884d586136a2007c7497634be4a45d6fa1846842be377eda7ec8d871b61070a7fe62994be3da6f6c81740a5d179cd4b2f095dfb653f3900714692bccd691637fe8e22bce103823c6ef9f349cab979154b9db740ba0e8bc5d21d4bba68bd94e0d23bb8c0d378e2cdbe013f73ddd95430778497ca9a0fe822c8cf4ff80afa8dc711cc617e781144cf59989a057be91f9cab5df16d66979b532f8628bb4e27934fa648ecd28cf51ff720072187bbdebc2cc32d0e303f31ce0d6f24ab8a11b21b043e73646ce3baec23c6e42efbde51220848af10c3b65b672a673e0004a24ce9317e0b33dc9d5d102698a94de5265dbbdce62f9582ff0349cc82c317093080b877876279c17724c07ec48cad56b1d4e19460becaafaabc05ca684af974c098b5131a22a8e65f199d3cc6549dd6fee8dd7eccc64e28719f7819dc6f45f9eb8236d238ee04a8ed030f462b62753f8c519310cbc1b673f7189f82ff1d87c4d0afe128866426b3ab903205da396f9e263e99ec93d7d1b5e273535ce8658400e402fc6a4bdac7d4a38f7a122805f041c9607717696762c58cd132f7006f8c2d6b1f9d1d4a80cb3811b0ff8d205d46b17bf96c2e9d4b346587dbcf31d25768fcc48e67c82ab58a779b6f6f5adf69b6195ed5e9f5c9e3a0de0adb727f5a6b32468ea156eae71c3eb9d6012383d17e4e0412e499f9fc407afd3dc2119e9ce567f4edb2ec69c0b65617189f9656319e1911e8e1cd3ff864ceab20e567de7ad8e82b0c2193a56127b53717a0378b6d11a940026dc5bd5c56af1221577597d33012816c30cf1a172e5803128e6fb715fc401d4448eb842d2f2aa064d3bccde97640c56e3238030c30b7dccca1e96922302138e9390c03b72330d3741b75efc142c59269f1526ee70732bc84c033b1b56ae801cb93233453ac22d1523868a57742b1cbb3fbed303e2e9b8cfef182d89b8f8a7af4798f4389610dccbf0725286c5f3caf9e74c512917483e9b8a60d5ec67edded1e78564218f6033dabff4147e2b17ddb84de0ce57dababc669af1de7d311411da36032dd3be8b6ad83126382aad271ca3b26159722dcfc0ea029383629444bf4c789b36d4e9a1003437199005bb63cb87bae2f3ae2c23e9888fd7f098b8b33a2f01f294ceea8b3162f73e5854d751b5b2b00cccfb3a4918a1e8ec89268ce04955775009265dc94b60663863a6f44051bbb89976b13b4e33e42f4817ebe95dcdcd0b7b6c4518e0476e5e9a3699bdbf3cdd38c073bcec13e4d57e91ed50d21a7695d1ac1dbe6b475fa9fcb1f86eff83613dc796389200aaf1692a509f6cd0428aae90ab50d9745b5b6ce16b6b8a25e143d02ed0cb90e9bbc6917b00f635698a4151f1e039c06661600b6fb4ea662786e834064cc86ebaa7e16145e8cc35da1635829308a339fb4f93e15e0bcc7a55e5883141f633f5c9c3b72c04a759d28d16e14c10beaca942fe0896951cc5e33ba74f856363fcfc69927d139fcf8300e23e5602a24425cef674557d8c77d2f86d2a6b68bbab9209f4f9b7cd5e2720ab4aa0086c5c7219744fc52330fd0064d558cdc2d6daa7908a12c016ab93b4c39c2f49db04906f4b21483d7f7b5ccd0cc3574ff2fac6dee2cef4c9e8c5e12d8675fbe4df116460c371fd824aee40d5cb7c754705c52aff952f3264175b36e0d0a92b29b6389c4312dc9d1daa318abbcf69a17c96c54ef724edd5e3c831d25240747910bc07f8f89e907377d41194e8434aa3eec237d81f0b81181f2eba289e5684bb98f773fa3cc3ef55de7b21161e19cefd2f514c7fc1700f6bffc9d4b9547cbb592baf99a60c7039cbe2ae5be3b6394115cc846eb29e7f15ce318d03e46949e4703f36b6caf8e96d5cc381f4dece230fa5eca683114816014b0052ef5c6ecc267c1b7e69facb6547127f38cc6ab8d5e1e043aa6c7ac35fca66797d9255c0769894c237c7fc8eb68324fea9b1f2a665528ccbf7c862ee384f5ab41acb50911499cdd53f7b68d61dcf050a195b0a11f6769b31d73a49e2fb8aeb1c0fa4895353443ff9cf34c8a78e6e8dc8a1e0620a4a8950db0a48dfdd34f3058df6fb6d5757fe3643b9003cd86b5e903ede7e00c3038515ba54b5c20736cad5f68c5fa48d95cc67dfdf12a704c8af5be9bca907086c9185b5be79e071b643f039e908898d122b30317ffa8de9e9d0d8ce200930e26652a8e43493af9238f529f42696dbe2624d6cd76f4db2a913b41880da4f2108a26b9c00ccef7931d2a5662ae48b0d5ffe8714d86f1112a65486b1768b078492c8df0091d317be1c367b30e0577b5bc349888da90c8b84f1726fdfb6b0e765ec12ac491102ba07f3e13324426cc59e70685e8c7a641285346f99c81a9a05953c0b9324b0ce5977e5d8db68c694e892bbed15aa9aa21a4d45322580cc48028ff85c581b99aa2747f7eb7ff0403882918071c4a71ed25d076c983eef52f30705cc5e1cfe08c2f1181f2d0b922bb03f9958fa81a2bf8eeb816d86a79ee668cc1b42f0b391d90bce53b185c4f912787c954ccddf89e50dc93dfed4e315af6895dbe575f9c4712ddb83e30300eb37317c842118e4a4c2a6d6ba23cfecff8d6019d1fa635dcafb8b9b52110ee61590143cccc65bdf73ad0adb6dfce69e744063d5c3049a9791aa5d6d1e83f29b84e38d599fa9743f6210d96cb841056fd1e11fc5c5b1c90d3fd2b304663d0ab6fd14ae200d874e1f36ccdb29aa49726c21f3d49379800d20a50511166d2bf33089107fab2a0461d75539042021a6d379b50d771685b02e7e0dca88a55e205433ad841cdc334d641710eeeb06d253d4c2d94d3789cce8dd52739e7e89558b943145677e4005ea4c7ce5bc72dcd3da36b1bfc4d4ec2b6ac583d8a18d8c3850ab5caa7f4684ab1ce854f281716df4e271083cea5969464f2b7fc0f8f2743b86c4b739ac847ca4bf8a985efc3287b057ffcb76cb314a5544b13771aab163538fad39789112e9ded835600c337bbc9db74b02275731f4017d563c9571ce981d575f8d855ffa1a455e564a40b43971b2a5b1efd9068332f5c25d3de5740a9b9a88917ca183c3e357dc7a28df103b1d36d57c65864e1693c0f5fafb5e60333e14cfc50585c9405475079ab4ba217772982da305da7b564960bb6441d1054254e533e06c215dbb367bae77516bbcf8ef45e7a46cf97d8d84b8d8fe748174ffd60227673ac028e06f1fec8453ef80840a3bb55653cb78b43be587ac852fd7b1519ddc88daaa867042adccffc9a34cb17bfe90c6847f101b705d4b773cb56f3be5e326300a762073d148fb1468b1ef1b582f34a9a10902effc851a074dacf4b1eab1635d408c571369e8202bffe84d0b9c40ea3fdcf1fd0f388f193435fa48c4e232f1768b6722aca39396e988b312cd6cf3ca8b7a055555fa05a277afbae7edadae8e43042f825d8416770aba14ce5dca782fb05e57013b3c4a5af3fb0b8bb8258465649b3ae621de137cdcd0eba82999b2e01457cc010912edded5e178e18a91b66369c06c1de61a8778a7743c8d07091c31ded8f7d76c1e1cee7bbbff8ccf5ac168db05c845911aec1358b8f8a741cb78bf5f721f5978d30ed8ac950623a9ac140bdceb18d9ac3f587a6e8b9c4b82f9fa1fcf11b5fd27610a00a8996666ee4b6fab7d969bb2617da6e6e498701bf2a601e229f45c0d30558ecc620fa8cb257bf629a970fa6fa31acd79aa534bda2ce1cbad3ffc3c289a4aa68b5c98534dc2cac21d2ebfd6515a27e93cd29ae537b75f51401ed273933336bafceac9a5dfa2110a274fb9bb82f1733800ed036de76934d8ad3383f710078bb3ef34bce4e9cc7a6849bd3d8dbac9a1f6f55c2ff31925852860607f8fbc4451cedc79b61200e25d5a51a0417d8b04f25d4208343c4e5b49942ddab7cdc563f0bcf2baa575aea33370228a168da4f5d88828c573d9cc580d5b4aadf9e099a66c5ae7a427103e00449ac5833e048661f10f6926c2332b3f74fc5eab76924138a6c96ccc32b0c692db4c3cbcd9e11d9b4918ccbebda5ea567c3f6efb77ec77f799f4d6c5ecdcaff851b23c6cf4c6e602cbb0d3f4f9aefd397214da7efaff7382cb2fea89d6cf6750f871907e0ddc26508271c506cc71d8852113a7934ea688b01af6e5a2731556d775c024c9aa1b747859da8b7e9252412d553da4b0039e2a3df7d32f3271a9db28050d28780fdf81a9a03cf3007b85696c64a5e297d005edb0cef204322ec767c68b2f2dd68201cc1398a4321aae12eb241aaa9d3400a1fa155b5eea4a592b5635830c8177f377d0cd971d89ac868cbd89258abb5e828adda9aa678d53dd144125e76f592518cf3457b07b2d7e69128e61c1dbc0914010b66a91bb46dc9fcc43b4caec07103e817ee521d7ddeb8d90ae40ab7e805c31119934093b140c7ead2161c8e32e4074b2c7a852ce9ea55a88fbfe7418ecfaa586d723d5a3d99186b4a93a2967890d3d24f6d2934bef70211d4cf32be9c82f7e68843c72ecea766316ce87fa837a3846013e17190e61fbe70cb0d0c3d565086ee6dbb184055f65b8be2ff461da6bed78c259fe5c7c5791241fa78d33116c76a50bd2013f9011295b0cd8e93b6113b09a7c893f8fda3673a4160ee9ad05e22a0e580b70892869e838da6d58dfd674d39bd90ed5309cb1ca208d0f977693f604be86ff85cf4c4c33c0d8acd001f4e74f459e8c0fd9734680947fb248d8792897cd49e3de9e4816f1ed49f9e50bb592eb1aab4fe4db712ecb64f236dee73aefa81d6fdafd06f5b4af36945858046b9ae02ea769e2cf1826a41612a2c33c1f1d2629fe3e990e9a967edf3756ad940941084b62b3328f1988960d8c394db56438815d19a4cdfb949ed30f12eb03a2e4faf0965e2498090a6b3e51fa49e15525dbcc4b455af263c5f52d453aae8254928c0a144a2e4235482ccaf1807cf66bf8982fdad66a945d04e30bc3ec35268ebf4bb21d245baa183cc7c76553ca4bfbd903fa0eeeed9dd8d7ef5eb7aef68afabeadbe7a386c59a9b573db60bd3c19438b69b7033c33ef85bc228c59a9b5d0957725f3769538b51d1dc3506e845b4b693809cb75716f70209cd0af3bb42ce1e197a1b7a7c75cb82009a3650a996f3affd5c0bef0a974d5b24ab6a31f999028df900dce1cc4c07654cbe9eccf2cff3f9d4f9294e3bbdd59187767159bb9409c00b198f9daa6b8fef063eab77f0986ba8b9fc46ee44b58290546144c3df877bef7e9568e589a13e60fa8be44886548472023fbcf372c1c22c31fc74ba232322c2ae9bed4b12a58010189bbb9a2e727caf0fb641bc2ec1fd494ee0c18b0f16fd05bac9dbe758653848dd3eef72886b4fc33cbf2ecc867065f9207c3e39057417284ff8ef0c87abb13f52176902fe3f12a3cb7cc8ae20eceb05fc857b52b9a28b7ed2ca16fc012f474e8034be56e33fc4a0f573a7af73f5f2b4ce1ad5f45b3269a021ffe44abd759c4cbe2b4081565a86c4cbd83a0e92916bc98634e81f1b983871229e14ddc7b336f6f9fa0abe013d70195d6979ed13ca3cee44062e6354b04ae2bf51d3716db6e2e2db79767d2e9447d9e257125fd429be30363f33fe77adb3421661161c44d2d89f7f72fd97356cd3dd2e5444852bd3a7076053363665e9586e267591f40bd61ab03438a99a61483d601b83431e44d78e0611c49db2465aa3302b160a33c31001461eb02e4550108a04a9660dbd38e7233c0b6d00103b722828da2bcbde0b83b38c5b04c436820c3ce2f1f0b70c358586895126e0e0d07888dedddb5f7e6d5517873c4d451b41f9265b257c7f44dca6ca5539b53cc9b674a7656668e60239ac21bd92a42800553f6f3dbd6c478c33654d243b3146325bb4a4a6bcf4512908c850a3ebd7d78cbb39f7c80e909f9ba8048c466307a70ba63d8dcc2b9518f9b09de0a24e64dff8bf73ab85348b6b570150fadc89f0586970ee18f57d7b44632bc125d48c0ef379c93c20a11910745a3e6c3c4e63a9e9abe938f90193ccfcdc592733b527534b58d2e41ced9cc8741d3cc52bd6a2cf64e0cca66d6904fdaf9cae214f588152848e5a629b2bc580bc2e9fd9cc41e45a50e5711fa74fe603a5a0aa1a38aae9255f17f44260bcce1535f401477bfd331616cfbda23baca010924a7417bd3da1c31d03136781745edc0fef932429c71d1fd778184501f6d6d06370bcba33b9f8b63275dcea671b38a7a87378b9e7cbdb3e4a40d42a0d228fb640cf56ec712f8f20ad12023eefe83c7a0c6e7eb65e69d64422bbb682e395aaef8f2fad14da06e88c255294c965f68b8c7c257a7d4ff3760c20603da355485f119cf5b495bdfcd576d628f5dd58c814ee0884d1fe15a34b3368c45a9caa7cb8debc2a3559d54430a1a106420036e7812d5e8dbe25988709d0de21e56322ba3fd6303771f72039420f2f441d77e1f61ca93da852638faa3129508fcf54adc33c1dacbe063b9e8491f2a13cab4e00a743ccc59adadc5099f9d0dbdc856a1f8ed50e1bd7f877ec8c8e5c8c36e19daa787c89e9533332e26b27cf57cfd634f5ba40fc7bb6222bf0d57d334d3c599fab2e0114c8d89e15c4fe5d6fa13d6053dabfae286eb777b7dafc7ea30027c597092ddfc584bd5c7992b729f0aa22b7d1d69abd44b8f300105841b8a204e6ea76243503c8428366e55168cf875c63beb68114390a433c531340bc4e9c860c8316a3012d9b98e6d0046dce7c45e80c9cf93e261a62a4d939636e5125e76c465c2a30d439508689a89de31c085823bf4eb7225478c951388e4b0556aadc7db9938b4d8b6e7505c9f115623b21e6901d1e37aea1e4c8b0e687ed31112856211c5bc4e24fe274d86ffb65146188eb7bc7cd24f13510529ac52e4ac8816e00cd287f2c51d1c2ac394ff53fd1944225ab426b8ab95d326b08ba53dc94ff3897a703225edff4ce281f2f9ad1e99c84defcf57b6153d5ffb8cc742398f876dd8133f4f86472ecda5e67c5849687d2bf3614fe85e49c3d75474954980aeca4a8d050d169988ddaf42baa41ec4fe01990a8eea8c7319620f09d6ae34c2c956f369173e30fb30f84d563a668eed12caac6b2a5b12ebbb144978e35948a95dc987a53fc008fd8bfcc001dfea60231e3fba75cd8b2fcd8efea787f4f0dedfd8d6867d2dc24cdbcf820842dff041027dabe6a81d4bfaa40786a9e868748d94650c0db4a053ef499145c6b427c7cdac31c9ccfc4c5be11fdcbb450f4b4b72fe40373692c9964181aff085d2a584a0e29543b7e99f0cc9319560da89032d38da5f492467b3abddf852040d1a177db1635770e5ca8266cbd82141200cf9f86026c86b44e0c851c7785d5598a3fa580a19eac6fef98f6dc1f34f3b81e12fb559bb26d553187dc7580b68430b8e40f592a1d4c7488a14cd4aa850b1f8b4100581d5a0bbec94c0ea225dea58e2c782758b3dd1dc66029a871b82b5f3f540b97280b4f26fdee5b11ad7da90b8585f90e541e46a548d9462a0d8526caabd1a705438ef7cc0ab79df1234c0f5df34188824da1aba09352e295d6beda82ed820ec5da9b0671897c7cf58be9996f96bdaea8881e108c5f49198c1179a9156166333cf388473c1e7d157f14ff81d68f19946c9aa775e93534e83a77173f0441b7042e4e28fab8f29c928407dc9017a3c1280dfadcb9db48ab1c6495644dc76015d92936e508c48cfd46c18903ba1d46ddddb8579a255721cc2ec08d4b771c55f438d4177d5df93e59a77f665a14f80f85e2c6cfb74c7b8940413be3bdfd85f8b4a22a3c8e54be415492f76fb412a1af6420cf500fabe8a05a285904be3bb4916fe82cc96b8b0f6a38aa88f6884261ba90e2838bbc835adacc985e639f471e46f71f083f9652b20aa0480d0dc2359e80a353017cd9bb9db1c58402db7652c39fe8e95dfba6caad85d1029e58d81af1ec07a1dded68d0275241ff9b1da740fc588ad4dbd87346759163b4ce774fbcbca364e8ed3f72e7e9c5bd618773cc512ad671f54af835a53a293bdbce761c62b1f0f4f4ce7bbdf1f741c30411f053cbe77071029d034d213f5517af503c89112a0394c2efbca77760925bce4f42dcad45e53524d04f9396c455f41943b0b54223f3ad1aa09644e94104fc01e8a18d104baa973e9db4f48cf89d63718f58e2b0c50ce942e37ade89b5b1979c2a9ec6cd575b32870ef4ff143db615a8898374a3d0ee75224cf10b8a3525657f8631907068cb5c1201eb8e3ab1e85c1109cd27dae1fc2e0dc7d03fa751408ed83f79b96a83aecc51abe57320e37849ad69e7fec4d500357e6ceed0a3ef99f0b347c7be90400d80318f10a6fff3b32becf107558de3908fe2f5bb768bc0384fe2bcaf625cec7c29e0e498ebbe0fa20c400aa1488d38c15eb9fd6d720870c53ba61edb1e37e47fb5e8488cd3f042892287e1b0e6c817bd4e5e6934b704a5e7f9ae0e0323fee2ddc235c5ff4dba95a3d3100714c97a5c51db97240b65a2fed8404dccad67ee536769a636a279227cbb7981c5563bb175526f2d3ab6651e4022226478bcb054ffbd66d55c7d5a875a8f0a9887f76dec013cc86cb05c5bdcfb781bf9eeb6c77068abc6901b4b1811e1cf672073012fb1d1f10b89933e6e624d5edcf537abb514c04d1cb5d32443a6a1aab6bb93dacef15342932fdcc6d327e487503fce1772ee0a4db0389aa28b86d09b5ab5139b92f46105f410d30b79cdc4c34edf342c7251e304c2083a65f551ba287e8de97fccaf7ed5a1dff74193509398f2d33f88a757c302fe75b13dfc7f73436f7121b5382f9bdf551edb8bc0f10c5008628269c778705c8c50c46571f99466298952c78e11cb0da4bbc2a906c112172b680a27727022145967ac8e759f586aa9114c0ab51187a3f09ff5597a88bfbd76c3c328fa2d896ae226c473861b3f62b21ff6be4fe6e062e7e7d4010bad7a58e300f145e8382414f8e64ddfe297488bc9ffa06604984406e5ac5ab1230976c6c099a12972f299f2cedb738504951c07f219a7f88a62ab27322272c182f4e961709554e50114931ed2fc57b3807f950d49a9dad482d56caabd8c15f8bbf96483c43351291fcd416713de107095cc8f559a5b5121ca22cfb08272916061b5a873de03d7da748d1556f6e3dac3d67daf91fc71833a4f07739c458c1088d563335d51f25f169479be35684d9ead62e4ac43ddfa4b87c1caccfe0b9c1c9569b914623314a121b14863f64c150a07a56beb67e0edef268b3952884481a83a6c756bac845ae171020d91ba189beeb5aa3b1477c6b69c4bc41f4e7ad504c0ed75abaf67c7a228868ce751b2fb40b0aa4f092c696f83f06cf79013a03f18d521f265caaa0a1ea4b9f5ebabc771d2d8ac6070c27b7c7d6b8d8fbe922a63de4e38363022534b1876a8aea9bdddc3db46317ada5b8277d35fde33f86a1c571bbc99aed64c2819e8e944c410ed8dcaf4252070ba9e368daa9d3a7f7505b6e37c89f62fb2fae77422ccfef61aad75b6a7215023639266c22b70c88abf191e7dc2491e313f91a3714a685c2cbd530e63db36dc94ab0740e7f93f6c3471db96563ff693a70b35c1d2e33f44dc70ac8f031b2e1d54c22ebef38e76988a8a95c5762416c42d1c2d6467716f67156f813c308e142be1c54ec087d030f9d336990d2c534695b3b7e3e1660895cd504597e22944d4337d3cc5b97ea4b5dbeccdb4bb7b763da5748f307f36b6c03950a4fdfdbff6cb524b1ceab5c47104b4a2e30b854bb3f339823c7afe949b0324aaaa41cbdf333939e4c494f67d1b460655de5e0ae88fa0d94fb59f73f280d792791ed2a640a1ed0a520f974c2ca7f7c96e6f3106079125027aa5c2c42a14e7b8140dd5ac1bac84875d4349376839ce046fe96c813b04ae426e69d7f2f28291a7dde117de32a6d17b5d06ffabc8ad30ae6bbd16af2a1d6f8c445daeebf8fc411705efb16ab08ce80d480bcc0b5797549343ae6d7f37043571e684cb7a8543a07e7afd2d5ccef29d312db2524f2cb881f172f1d5ae860238a2b76b3b7960a8a6f77de891c8f456ae27a57e6d5d3e853e3abc5edebe281fdfe3e71b8ac19354856d17169fd50e4a6b456790f65e3e40afba8380b901de67540328e72820bc0e98ee5257871212e03f78724dc58abb889eb169d2926f0521e3056aeb0921c31122e205765461c112f50d02006978ecff5fa5433854cfa1e3f3af02f92505ae8d81e845bd6e2b404b1264fdde9f5db687bcb81a7e37aa3783352984237d74a262747ed2a38c0b654f77b3fe2044434b5c53ff486f156510254fa97e809cbc5471629e034e0256ef6f87dc043f07e6026a4780bb0a254e644cd9c951cfd3fa0066e78bae2cde51c4e39b10f904aa8d8838e785b69ef96adf7fdd7d1c997f608805f031cd3f665b8d2b715f066dd62ec010a59b4a86944a7d82674aaab5d3e16f8e7228cf1ae80a7b3bc9f2abc8326ad0550ca4dea1a4055f083de39f6f484de6c0be5638aa901cae562f304ee23836adf199d11a633ca7e3ae19d758aeb78c1b376b0fd4b58e6b461abf117261f448d1da375e670621db90661da164699213c01406c32e68764a564ff5f941d6e6b9587c1e9b05219ee8f7eedaa81b5810172f685c34ac5417eb34fe7f75fa6ebf66a6fa7ff373c3eeb982cf1a3044bbc5d3f1078a2ef7807c8b5079918c950b5b476c80edd0a5b4f5bf788046308cf1e43c3bfce641989030e96960f9bba6c53dfe7dc39e0e4095edb0f3f9bf123c29ba7192ab69155260f2fbe9c715bdf6e98d3f099e12db70f92e6a460fb52936afb784cc0a762ff5960e576f58cb09dc1ad75fe8d861bd132ffbe40b9961de55be7f5c451f01ff100105c74a2aaf9d9f96cde45024bbee97a59a5fe85e2d6dcf8a7e466101761cabf5b8e7ede83ef01a38f9894b478bc8ca5ee84a0d0d3e0ccddb3ca408de72f25809e73f5613f47a3d475a511c5daf4d2049454efb265d6910ffd2377e3dc248c59cc81ae312d36c49328706a1b55f16af52263bdff7620e97e2b83ee9d871818aa47b748a319ee6fea0e60b79cf57ea1bb4f190d3cdfbb0638532e41df0e405c0376b6cd80aaa6fc8b952bb14d2b34256d02b274e9149de87c79625b3ee4bbf8b71a42f2318b813913be309e8830b50fe33ef1576601c238292dd665ffe64c39f5c344b57e41c21705d1a35fe42250ed34c69fcc6afb204a18207ab1505c0f7634164b9cc684adc2fc48b9db1c8d29dd9255a52a8bfae6f7937d0c2c944fdb804c2714dfae871449bf602e8a6e3bcaaa3f3d3a98e50a0b1c62918ac162d75901cc6079f6250ca0960dc59959bbc1d5fc765f19f5d774ee5033818936f92d91b0406d89aab92b5b4772d8031a08935cc12a92ace2c304bb34dbbff1d777762e70a98fb3f7b4009b21a2816abf13ecac827578bbe39e9c571abead574e3055b89cbeffe47cd1954a815b9d0303164d9fe52d92f558b11a8131522014d681d68bbd91ce8ee0e1130c2d7a53eb81c46aa1768f11410049761858d702377f56a7e129c48218e4d2b6c744186617b6e574c635b5038158e8e449194f7a7b3f40fa15b6d5c1aeec5c7313773302243c4d78e90c13e5ae4a8ac80d05e9c3d2797478bcc84ae14cecefe7c6c094a6bab03272be5b79df937d5aa279f0886829af55fab1fe8dd9c959acd082dffadb227eedcb0f51ad3900666dc5f46d909cb0c1a5fd81e9b25fa26a96a42a164c118617646d3b0b0373ffe619cc4129ded90d28cf924e753fd9b908aa9745b731efbd46dd8202d9faf9b452ee3f648afc8c4d74d41e4e930bf858ced68669a16ae30de0826ee9e7972f1712669f8b095277c7f77377f73a8ebed9b1545726710805f6dd6ce5319aa65210c489321979f4eacdf4ab59af1f253a83107051c6049708553b22bfc1a1fc4f294b1ec1054e34bc31b27825f2793874cc4dca34e869bbf2d40710f1bf78059848d9e34b44a5d227dc39b97118d5476d6034135cea38b293fcd292ddaf4f8597e6f597393a292d1879dd390c382c24ddb405d0a8fbe6970963622757e060c345519dae5eb81051684b9159c74f710839cfc837675d9e77e70a63bd24fb8e557e3f19f89cb9322eddfdbdd5ad3edec8ed1ef99388aa8ad842725163b0f06f3c1628e660fea7def99da4a6f3ca509456ef0e099d5a705abfff2cf589debf99d12a7b718112dfea7e37b97dbaac0eb7a3dfd9cf495ca0d7b2f69f439b393c0528836c739f525c109050c61fc4577026d7575b6ebf32685b8ff730021a18e0333ab64d369493ef42886cbf1528127e2e3ec695dae05d76187d346e5df46cea02786478ee387a85978e8bb4f0bb50057ca5edf341b79437a71ee0c62dc1d8e55e59d66b431e29450fca2c18837ca0dfdbc892694c2e5e5c71497536e250484eb4947603a40fa697292e2b3191105446b000e7ec2359bf85bbf3f46508d2046456d55b49855d96cae99f4b887e2c4d3914f4fb1cf8d40767bce0bd9a9a8a4a6be9ba3d144ef7eab56ccd62dbfec9c2535f05600579e55b83570d98242508ce6e5d24eb84fabc6d2bb3b5191d622e7aa6cb35383a715103b55849ebcf59b15b69f727e96f710326304319c2b5420fa14b4afc8d668ddd646a48d46f8f997bfd4985bfc525c93ccab2dbc3b093e12047cf09ead285c8c16ba156eb70dcb6190f6048ef79b928d02e99f7c302767f67566279af7fef4c2f33e3abf13bea5911073bd0843a3a1156f0f2cedf6adf3fe1d3b201a56e6023e9a2d85933d101818485c44c3462d01cdbc74c8cb738289ba697a77fea9279a56b1a2a1c80205faab2476184a318cb755bbd301f8f9e3a53aec72b4232bf205980935bc6be1e75be855aee364e01bffd9a5b2891d98d7bd0216d996a1ad9e04229a3499aba053c7def192fd195e78fe5ae2edde301699a78dd2020d3d9a1b980a8f3971857d9731ae7dcf4e857f9c123b11af3953291112b00dc7826bfd3604cb54f37e7f5b8fc579b8448fcda674548fd80c4fdda3136a711e05bb30c9225779cf7288a9fae2e965cdadee424c67c4965f06f9b98d74da305c2a852642df7930b08b6b7fd353c1780c896b506cf45f8819eca608e3da8b9c6f86b516824479df1fa63730541b4a61d3949a94b7d2853b224e9e4f89b5216c14ff06fe54d061ec94fb247cbc8a9cb9ff3fbeb76203918438043018290ebf09708be8e0e39b8ac45cb74cdf092a60f84faa04ec25d5dae6e0fb1d52071cc1a84c3a0b957370b4a6e292baec91db66a6cef5582a8df125869eeb6cc4d9a0553dd75d5910e3891035928cddb2bee6b00bfde57d78e91478db1cdffbf2bb5c135d4135990577e6edf122f99f25b7ec4180d8b3458c6591ee3f0c8f8d4459757c8ccbf6d5f8595c3337695d20496f83f9b8c8845a50a905ce1cdebbe6e5c029c1ac5408ff4a20205785993c79de7335de90076367213183bbe29389a85c6498a5d1603cb5a9b2c9136c4b65c5702693c4943c632907e48fb230f7634bb80ba4c7469e7ef89b78381d157509ad29be6")
	require.NoError(t, err)

	t.Logf("SPHINCS+ scheme parameters: %s", sphincsplus.SchemeName())
	t.Logf("priv key %x", privKey.Bytes())
	t.Logf("pub key %x", pubKey.Bytes())
	t.Logf("sig %x", signature)
	message := []byte("i am a message")
	require.True(t, pubKey.Verify(signature, message))

}
