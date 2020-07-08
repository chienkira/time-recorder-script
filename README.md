# time-recorder-script

## 利用に何が必要？

1. rubyが実行できる環境
Macの場合、デフォルトruby（2.xバージョン）入っているので、すぐにこのスクリプトを利用できる
2. 勤怠システムのログイン情報

## インストール手順を教えて！

1. スクリプトをダウンロード
```bash
git clone https://github.com/chienkira/time-recorder-script.git ~/.KOT_script
```

2. aliasを登録しておく
`~/.bash_profile`ファイルを開いて、以下の内容を入れておいてください。
```
alias kot_in="ruby ~/.KOT_script/kot.rb --user=your_id --pass=your_password --action=in"
alias kot_out="ruby ~/.KOT_script/kot.rb --user=your_id --pass=your_password --action=out"
```

3. `~/.bash_profile`をリーロード
```bash
source ~/.bash_profile
```

## 使い方を教えて！

ターミナルのどこでも以下のコマンド1個で打刻できちゃいます！

```bash
# To Clock in - 出勤打刻したい場合
kot_in

# To Clock out - 退勤打刻したい場合
kot_out
```
