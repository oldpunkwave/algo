package countcommowordswithoneoccurrence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}

	testCases := []struct {
		desc     string
		args     args
		expected int
	}{
		{
			desc: "Case#1",
			args: args{
				words1: []string{"leetcode", "is", "amazing", "as", "is"},
				words2: []string{"amazing", "leetcode", "is"},
			},
			expected: 2,
		},
		{
			desc: "Case#2",
			args: args{
				words1: []string{"b", "bb", "bbb"},
				words2: []string{"a", "aa", "aaa"},
			},
			expected: 0,
		},
		{
			desc: "Case#3",
			args: args{
				words1: []string{"a", "ab"},
				words2: []string{"a", "a", "a", "ab"},
			},
			expected: 1,
		},
		{
			desc: "Case#4",
			args: args{
				words1: []string{"bjxzvssdoq","oom","lxrrvf","aoeselhvrnw","awnornqyztqlza","bjxqkapuvaw","wibxruerngdzgjd","rezrwdzvllpbjpnikhzraz","pswmnrsepudx","nlicjldpeia","glg","nllxfcjjitmsuugmr","cl","pysmpgjakkjnusfopphb","zxlwcdjpn","xktsfnchwrdesnf","qptnoxxgrjmvr","exlfwjfsbsirbbkyqjtinrrwuhh","rqbnghajxygilgdjejopyuwyjqrx","vrjkqsicuqoalqyaxkaaogxbf","ixnlltqbpygmpjuspom","izajsxotcbhzdnkujwgdzo","b","lighabre","i","ljqqbfddipvcooh","hboedpepeeunx","bkhzhiefammwqkhvampokd","ptlozguwmyyp","loeshsjgazzwvs","kyrltbdzlymjxtvwiiq","fk","mbjpgwsahkgkehlcoqbhunqchxj","nfyuvlrmiturheb","cyqwsiysmoirurj","sciqruywy","podsrhmsozan","nlyadkrxhdbup","gdugldwghzt","wpjm","gjobdekmjisjgadkwwemnmco","dkjdtimdghvlhuetxyaklk","iwqylhdwqbwaqdouowoodhs","mn"},
				words2: []string{"eeormvovhzslwsqgzthlgntgzc","zfwownznh","suxrkdbjdjjtbkjucsbyk","u","y","lbjooktoctgwbbptiffytquha","dcsxrghgpultkatbecjadbespvww","vwduylshcpaiu","rtcxwctvquaiuwkgvdx","a","szearxmdqcismljmihbtkcirztdnrc","htgmuxtxdunsvfizb","hybe","nsegkgwcvopncmfpaahhhjeuqjosv","jtarnnpppxtzmopixeijqqahkd","hazcgrrnpourkyoeanodejiptne","kurhokvhixihe","ljwycewmecfqdhtxiokjn","qgjzzvpyvwetlsvcsw","aunns","nwcnfrzzvxafkfjfnczummtubikji","nipiygnvlfntgpxfedj","mgnt","xvjehufvaqouhztnmts","sjtbrfjwtqxakqktxjaljrbwfoxvz","dfeujeikfrtrpiafrgxvjlkpxtog","u","ggbcxoasodaqaazulrxjleecexey","inedrgssajhpygfvozigohis","pevxwgfzxebybfe","cgy","fnhvlx","vxfybaebkoq","xvhx","mxbqjtanctljewwjjlbpkgbtsm","mlwagamcikbcpuexhikmizp","qeiomipvsoqlsnhylulirrcwtqga","bwemqcgyusuauwlpbjjru","iimcbidtndh","lpjejlkmxtlbyvnscy","dlspriicnyykdsyvswlgktavwloq","dib","qoptbduulgqwquvhdvmwdz","xrtxghrbfrhpzduxeljnctgg","schmbsaupayqnpchn","kah","itotymryqufqpozrwmvsl","gurb","xsyocxcmwvqmnmxthfemmu","pkfdutm","hkbwxwjxyuld","ukdqszfjckdunnhpevw","kqfwytdvnvjrchiwprcqkfntqticsc","zjmsfwjddgjiypsmagdrujb","gn","ebncbjvhpbjecbrizdpv","nbfehcktwswih","sttmqcdmobdgtgkyxydyovahknjbsn","sryyufrtocf","eiicpwknxrzqylqpybhfd","pey","njimttradeoa","wcogjdfr","prva","tyxdmxgw","wluzocppg","kzm","wbyyperlkflaoxyxftzwxvmemof","snzpclbulddnmmjmpdurcybo","mowxgpmzojtmympmt","uvtnojjahrovzmlukf","sykhtgejlmbzshhneoyqr","ib","haqkkizidifykwijm","csjtexnr","yvgr","vzcxbtlthrynnawxnkxdptp","yvxrmscsckv"},
			},
			expected: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := Solution(tC.args.words1, tC.args.words2)
			assert.Equal(t, tC.expected, actual)
		})
	}
}
