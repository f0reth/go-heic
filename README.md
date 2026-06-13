# go-heic

## 実行に必要なファイル

このライブラリは [purego](https://github.com/ebitengine/purego) を使って libheif の共有ライブラリを動的に読み込みます。事前に以下のファイルをインストールしてください。

| OS | 必要なファイル |
|---|---|
| Windows | `libheif.dll`（または `heif.dll`） |
| Linux | `libheif.so` |
| macOS | `libheif.dylib` |

## 使い方

### インポート

```go
import "github.com/f0reth/go-heic"
```

### デコード（HEIC → image.Image）

```go
f, _ := os.Open("input.heic")
defer f.Close()

img, err := heic.Decode(f)
```

### 設定のみのデコード（サイズや色モデルの取得）

```go
f, _ := os.Open("input.heic")
defer f.Close()

cfg, err := heic.DecodeConfig(f)
// cfg.Width, cfg.Height, cfg.ColorModel
```

### image パッケージ経由のデコード

`import _ "github.com/f0reth/go-heic"` のように副作用インポートすると、
`image.Decode` / `image.DecodeConfig` で HEIC を自動判別してデコードできます。

```go
img, format, err := image.Decode(f) // format == "heic"
```

### ライブラリの読み込み確認

```go
if err := heic.Dynamic(); err != nil {
    log.Fatal("共有ライブラリが見つかりません:", err)
}
```
