package main

//设计一个使用单词列表进行初始化的数据结构，单词列表中的单词 互不相同 。 如果给出一个单词，请判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单
//词存在于你构建的字典中。
//
// 实现 MagicDictionary 类：
//
//
// MagicDictionary() 初始化对象
// void buildDict(String[] dictionary) 使用字符串数组 dictionary 设定该数据结构，dictionary 中的字
//符串互不相同
// bool search(String searchWord) 给定一个字符串 searchWord ，判定能否只将字符串中 一个 字母换成另一个字母，使得
//所形成的新字符串能够与字典中的任一字符串匹配。如果可以，返回 true ；否则，返回 false 。
//
// 示例：
//
//输入
//["MagicDictionary", "buildDict", "search", "search", "search", "search"]
//[[], [["hello", "leetcode"]], ["hello"], ["hhllo"], ["hell"], ["leetcoded"]]
//输出
//[null, null, false, true, false, false]
//
//解释
//MagicDictionary magicDictionary = new MagicDictionary();
//magicDictionary.buildDict(["hello", "leetcode"]);
//magicDictionary.search("hello"); // 返回 False
//magicDictionary.search("hhllo"); // 将第二个 'h' 替换为 'e' 可以匹配 "hello" ，所以返回 True
//magicDictionary.search("hell"); // 返回 False
//magicDictionary.search("leetcoded"); // 返回 False

// 提示：

// 1 <= dictionary.length <= 100
// 1 <= dictionary[i].length <= 100
// dictionary[i] 仅由小写英文字母组成
// dictionary 中的所有字符串 互不相同
// 1 <= searchWord.length <= 100
// searchWord 仅由小写英文字母组成
// buildDict 仅在 search 之前调用一次
// 最多调用 100 次 search
//
// Related Topics 深度优先搜索 设计 字典树 哈希表 字符串 👍 248 👎 0

// 方法一：字典树+回溯DFS算法优化
type node struct {
	Next  [26]*node
	IsEnd bool
}

type MagicDictionary struct {
	Trie *node
}

func Constructor() MagicDictionary {
	return MagicDictionary{Trie: new(node)}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, str := range dictionary {
		this.insert(str)
	}
}

func (this *MagicDictionary) insert(str string) bool {
	n := this.Trie
	for _, s := range str {
		index := s - 'a'
		if n.Next[index] == nil {
			n.Next[index] = new(node)
		}

		n = n.Next[index]
	}

	n.IsEnd = true
	return true
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return dfs(this.Trie, searchWord, false)
}

func dfs(n *node, searchWord string, modified bool) bool {
	if searchWord == "" {
		return modified && n.IsEnd
	}

	index := searchWord[0] - 'a'
	if n.Next[index] != nil && dfs(n.Next[index], searchWord[1:], modified) {
		return true
	}

	if !modified {
		for i, n2 := range n.Next {
			if i != int(index) && n2 != nil && dfs(n2, searchWord[1:], true) {
				return true
			}
		}
	}

	return false
}

// 方法二：使用字符串数组每次进行比对，暴力解法

type MagicDictionary2 []string

func Constructor2() MagicDictionary2 {
	return MagicDictionary2{}
}

func (d *MagicDictionary2) BuildDict2(dictionary []string) {
	*d = dictionary
}

func (d *MagicDictionary2) Search2(searchWord string) bool {
f:
	for _, s := range *d {
		if len(searchWord) != len(s) {
			continue
		}

		var diff bool
		for j := 0; j < len(s); j++ {
			if s[j] != searchWord[j] {
				if diff {
					continue f
				}

				diff = true
			}
		}

		if diff {
			return true
		}
	}

	return false
}
