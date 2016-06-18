# vim插件显示树形目录插件NERDTree安装
### 说明：spf13-vim是Vim插件与配置的一个发行版本，包含了一整套精心挑选的Vim插件，采用Vundle进行插件管理
## 配置文件
	~/.vimrc.local               #个性化配置文件
	~/.vimrc.bundles.local       #本地bundle配置文件        
	~/.vimrc.before.local        #早于spf13-vim加载的个性化配置文件
## 安装
	方法一：curl http://j.mp/spf13-vim3 -L -o - | sh
	方法二： 通过 bootstrap.sh 文件安装，bootstrap.sh
## 配置(必须修改)
	cd ~/
	vim .vimrc
	在 .vimrc 文件总搜索 NERDTree关键字的配置，类似如下配置：
	 "NERDTree{
        map <F1> :NERDTreeToggle<CR>
        map <C-F1> :NERDTreeFind<CR>
		 aotucmd aotucmd * aotucmd " 打卡vim时自动开启nerdtree插件
        let NERDTreeChDirMode=2  "选中root即设置为当前目录
        let NERDTreeQuitOnOpen=0 "打开文件时开启树
        let NERDTreeShowBookmarks=1 "显示书签
        let NERDTreeMinimalUI=1 "不显示帮助面板
        let NERDTreeDirArrows=1 "目录箭头 1 显示箭头  0传统+-|号
    "}
## 我的配置
	" NERDTress File highlighting
	function! NERDTreeHighlightFile(extension, fg, bg, guifg, guibg)
	 exec 'autocmd filetype nerdtree highlight ' . a:extension .' ctermbg='. a:bg .' ctermfg='. a:fg .' guibg='. a:guibg .' guifg='. a:guifg
	 exec 'autocmd filetype nerdtree syn match ' . a:extension .' #^\s\+.*'. a:extension .'$#'
	endfunction
	
	call NERDTreeHighlightFile('jade', 'green', 'none', 'green', '#151515')
	call NERDTreeHighlightFile('ini', 'yellow', 'none', 'yellow', '#151515')
	call NERDTreeHighlightFile('md', 'blue', 'none', '#3366FF', '#151515')
	call NERDTreeHighlightFile('yml', 'yellow', 'none', 'yellow', '#151515')
	call NERDTreeHighlightFile('config', 'yellow', 'none', 'yellow', '#151515')
	call NERDTreeHighlightFile('conf', 'yellow', 'none', 'yellow', '#151515')
	call NERDTreeHighlightFile('json', 'yellow', 'none', 'yellow', '#151515')
	call NERDTreeHighlightFile('html', 'yellow', 'none', 'yellow', '#151515')
	call NERDTreeHighlightFile('styl', 'cyan', 'none', 'cyan', '#151515')
	call NERDTreeHighlightFile('css', 'cyan', 'none', 'cyan', '#151515')
	call NERDTreeHighlightFile('coffee', 'Red', 'none', 'red', '#151515')
	call NERDTreeHighlightFile('js', 'Red', 'none', '#ffa500', '#151515')
	call NERDTreeHighlightFile('php', 'Magenta', 'none', '#ff00ff', '#151515')
