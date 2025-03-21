# gof-in-go

「[増補改訂版 Java 言語で学ぶデザインパターン入門](https://www.sbcr.jp/product/4797327030/)」を基に Go 言語で GoF のデザインパターンを実装します。

README は ChatGPT で生成しています。

- [Creational patterns](#creational-patterns)
  - [Factory Method](#factory-method)
- [Behavioral patterns](#behavioral-patterns)
  - [Iterator](#iterator)
  - [Template Method](#template-method)
- [Structural patterns](#structural-patterns)
  - [Adapter](#adapter)

## Creational patterns

### Factory Method

#### 概要

Factory Method パターンは、インスタンス生成のためのインターフェースを定義し、具体的な生成処理をサブクラスに委ねるデザインパターンです。これにより、オブジェクトの生成過程をカプセル化し、柔軟な拡張が可能になります。

#### メリット

##### 依存関係の分離

Factory Method を使用することで、クライアントコードが具体的なクラスに依存せず、インターフェースや抽象クラスに依存する設計が可能になります。これにより、依存関係の制御が容易になり、変更に強いコードを実現できます。

##### 柔軟な拡張性

新しい種類のオブジェクトを追加する際に、既存のクライアントコードを変更することなく対応できます。新しいサブクラスを作成し、適切な Factory Method を実装するだけで、新たな機能を組み込めます。

##### コードの再利用性向上

オブジェクトの生成ロジックを一箇所にまとめることで、重複コードを削減し、可読性や保守性を向上させることができます。

##### テストの容易さ

インスタンス生成を抽象化することで、テスト時にモックやスタブを容易に差し替えることが可能になります。これにより、ユニットテストの実施がスムーズになります。

#### まとめ

Factory Method パターンを使用すると、オブジェクトの生成に関する責務をカプセル化し、依存関係を緩和することで、拡張性・保守性の高い設計が可能になります。特に、変更が頻繁に発生するプロジェクトや、多様な種類のオブジェクトを扱うシステムにおいて有効なパターンです。

## Behavioral patterns

### Iterator

#### 概要

Iterator パターンは、コレクション要素の走査方法をカプセル化し、一貫したインターフェースを提供するデザインパターンです。コレクションの内部構造に依存せずに要素を順番に処理できるようになります。

#### メリット

##### コレクションの内部構造に依存しない

Iterator を使用することで、配列、リスト、ツリーなど異なるデータ構造を統一された方法で走査できます。これにより、コレクションの変更があってもイテレーション処理の変更を最小限に抑えられます。

##### 柔軟なイテレーション方法の実装が可能

通常のループでは実装が難しい、逆順のイテレーションやフィルタリングを簡単に実装できます。カスタム Iterator を作成することで、必要なロジックを組み込めます。

##### 単一責任の原則 (SRP) に従った設計ができる

コレクションとイテレーション処理を分離することで、クラスの責務を明確にできます。コレクションがデータ管理に集中し、イテレーションは専用のクラスが担当するため、コードの可読性と保守性が向上します。

##### 遅延評価が可能

Iterator を使用すると、要素を必要なタイミングで逐次処理できるため、メモリ消費を抑えながら大規模データを扱えます。特にストリーム処理やデータベースクエリの最適化に役立ちます。

#### まとめ

Iterator パターンを導入することで、コレクションの実装に依存せずに統一された方法で要素を走査でき、コードの柔軟性と保守性を向上させられます。また、カスタム Iterator による多様な走査方法の実装や、遅延評価を活用することで、より効率的なデータ処理が可能になります。

### Template Method

#### 概要

Template Method パターンは、処理の骨組みを定義し、詳細な処理の実装をサブクラスに任せるデザインパターンです。このパターンを使用すると、アルゴリズムの構造を共通化し、部分的な処理をサブクラスでカスタマイズすることができます。

#### メリット

##### 再利用性の向上

Template Method パターンを使用すると、共通の処理の部分を親クラスに集約できるため、コードの重複を避けることができ、再利用性が向上します。

##### 柔軟な拡張

アルゴリズムの一部だけを変更したい場合、サブクラスで具体的な処理をオーバーライドすることで柔軟に拡張できます。親クラスの骨組みを変更せずに、サブクラスで必要な部分だけをカスタマイズ可能です。

##### 一貫性のある処理

親クラスでアルゴリズムの流れを定義するため、すべてのサブクラスで一貫性のある処理の流れが確保され、バグの発生を防ぐことができます。

#### まとめ

Template Method パターンを使用することで、コードの再利用性や拡張性を高めるとともに、アルゴリズムの一貫性を保つことができます。複数のサブクラスで同じ処理の流れを持ちながらも、必要な部分だけをカスタマイズできる柔軟性を提供します。

## Structural patterns

### Adapter

#### 概要

Adapter パターンは、互換性のないインターフェースを変換し、異なるクラス同士を接続できるようにするデザインパターンです。既存のコードを変更せずに、異なる API を統一的に扱えるようになります。

#### メリット

##### 既存コードの再利用が容易になる

異なるインターフェースを持つクラスを統一的に扱えるため、既存のライブラリやコンポーネントを修正せずに活用できます。特に、外部 API やレガシーシステムとの統合時に有用です。

##### クライアントコードの変更を最小限に抑えられる

Adapter を使用すると、クライアントコードは統一されたインターフェースを利用できるため、新しいシステムやコンポーネントの導入時にも影響を最小限に抑えられます。

##### 単一責任の原則 (SRP) に従った設計ができる

Adapter を導入することで、異なるインターフェース間の変換ロジックを分離し、本来のクラスが持つべき責務を明確にできます。これにより、コードの可読性と保守性が向上します。

##### 異なるライブラリやフレームワークの統合が容易になる

外部ライブラリやフレームワークが提供する API が異なっていても、Adapter を作成することで統一的に扱えます。これにより、ライブラリの切り替えやバージョンアップが容易になります。

#### まとめ

Adapter パターンを導入することで、異なるインターフェースを持つクラス間の互換性を確保し、既存コードを変更せずに再利用できるようになります。また、クライアントコードの変更を最小限に抑えつつ、異なるライブラリやフレームワークとの統合がスムーズに行えるようになります。

## README ChatGPT 出力プロンプト

```Markdown
xxxパターンを使用することによって何が嬉しいのかを教えてください。
READMEに追記したいので、以下の形式を使用してMarkdown形式で書いてください。

### xxx
#### 概要
#### メリット
##### {具体的なメリットの見出し}
#### まとめ
```
