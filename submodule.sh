#!/bin/bash
set -e

remove_submodules=("
git@github.com:senghoo/golang-design-pattern.git/pattern

")
## Blogs
submodules=("
git@github.com:pathbox/learning-go.git/submodules@master
git@github.com:foamdino/learning-go.git
git@github.com:radovskyb/go-packages.git
git@github.com:mkaz/working-with-go.git
git@github.com:Akagi201/learning-golang.git
git@github.com:pjh130/go.git
git@github.com:mycodesmells/golang-examples.git
git@github.com:liguoqinjim/go-labs.git
git@github.com:xiongyejun/08-go.git
git@github.com:ningskyer/Algorithms-Learning-With-Go.git   #算法
git@github.com:bingoohuang/go-utils.git
git@github.com:1851616111/util.git
git@github.com:tmrts/go-patterns.git/parttern
git@github.com:senghoo/golang-design-pattern.git/parttern

")


function parse() {
  url=${1%.git*}    # git@domain:username/filename
  _bd=${1##*.git} 
  [[ "$_bd" =~ "@" ]] && branch=${_bd#*@} || branch=master
  subdir=${_bd%%@*}
  subdir=${subdir##*/}
  [[ -z $subdir ]] && subdir=submodules

  # git@domain:username/filename.git/dirname@branchname
  if [[ "$1" =~ ^git@ ]] 
    then
      _uf=${url#*:}     # 从左到右截取到第一个":" username/filename
      user=${_uf%%/*}   # 贪婪模式 从右到左截取到最左端的 "/" username
      file=${_uf##*/}
    elif [[ "$1" =~ ^http ]] 
      then
      # http://domain/username/filename.git/branchname@dirname
      _uf=${url#*:}
      user=$(basename $(dirname $url))
      file=$(basename $url)
    else
      url="."
    fi
  # [[ "$branch" == *.git ]] && branch=master
  if [[ "$url" != "." ]] 
    then
    echo " "
    # echo $1
    echo url: $url
    # echo _uf: $_uf
    echo user: $user
    echo file: $file

    # echo _bd:  $_bd
    echo subdir: $subdir
    echo branch: $branch
  fi

}


### 2. 移除不需要的 submodules
for url_ext in $remove_submodules; do
  parse $url_ext
  [ -d  ./$subdir/$user ] && (echo $subdir/$user && git submodule deinit -f $subdir/$user && \
                           git rm --cached $subdir/$user && \
                           #  git config -f .gitmodules --remove-section submodule.$subdir/$user dirty && \
                           rm -rf $subdir/$user .git/modules/$subdir/$user )

done

############################# 3. submodules [submodules/user] #######################################

# echo " "
# url=$(dirname $1)
# tmp=$(basename $(dirname $url))        #git@github.com:paulirish
# tmp=${1#*:}    
# user=${tmp%%/*}  ## 从右到左截取到最左端的 "/" , 贪婪模式
# tmp2=${1#*/}
# file=${tmp2%%.git*}
# # file=$(basename $1 .git)             # maupassant-hexo
# branch=$(basename $1)

### 1. 添加新的 submodules
for u in $submodules; do
  parse $u
  [ "$url" == "." ] || [ -d ./$subdir/$user ] || git submodule add --force -b $branch $url ./$subdir/$user
done

echo ""
echo "[submodule] update"
# git submodule foreach git pull