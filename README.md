# gof-in-go

「[増補改訂版 Java 言語で学ぶデザインパターン入門](https://www.sbcr.jp/product/4797327030/)」を基に Go 言語で GoF のデザインパターンを実装します。

README は ChatGPT/Claude で生成しています。

- [Creational patterns](#creational-patterns)
  - [Factory Method](#factory-method)
  - [Singleton](#singleton)
  - [Prototype](#prototype)
  - [Builder](#builder)
  - [Abstruct Factory](#abstract-factory)
- [Behavioral patterns](#behavioral-patterns)
  - [Iterator](#iterator)
  - [Template Method](#template-method)
  - [Strategy](#strategy)
  - [Visitor](#visitor)
- [Structural patterns](#structural-patterns)
  - [Adapter](#adapter)
  - [Bridge](#bridge)
  - [Composite](#composite)
  - [Decorator](#decorator)
- [README ChatGPT/Claude 出力テキスト](#readme-chatgptclaude-出力プロンプト)

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

### Singleton

#### 概要

Singleton パターンは、クラスのインスタンスが常に 1 つだけであることを保証し、そのインスタンスへのグローバルなアクセス手段を提供するデザインパターンです。主に設定管理やリソース管理など、共有状態を持つオブジェクトに適用されます。

#### メリット

##### インスタンスの一意性を保証

Singleton パターンを使用することで、特定のクラスのインスタンスが 1 つだけであることを保証できます。これにより、状態の不整合や競合を防ぐことができます。

##### リソースの節約

同じオブジェクトを何度も生成せずに済むため、メモリや CPU の使用を抑えられます。特に、データベース接続やログ管理などの高コストなオブジェクトに有効です。

##### グローバルなアクセスが可能

どこからでも同じインスタンスにアクセスできるため、設定情報や共有リソースを扱う際に便利です。これにより、依存関係の管理が簡単になります。

#### まとめ

Singleton パターンを使用することで、特定のオブジェクトの一意性を保証し、リソースの節約やグローバルなアクセスを実現できます。ただし、過度な使用は依存関係を強める原因にもなるため、適切な場面で活用することが重要です。

### Prototype

#### 概要

Prototype パターンは、既存のオブジェクトをコピーして新しいオブジェクトを生成するデザインパターンです。主に、オブジェクトの生成コストを削減し、複雑な初期化処理を簡略化する目的で使用されます。

#### メリット

##### オブジェクト生成コストの削減

新しいオブジェクトをゼロから作成するのではなく、既存のオブジェクトを複製することで、インスタンスの生成コストを抑えることができます。特に、複雑な初期化処理が必要なオブジェクトに対して有効です。

##### 柔軟なオブジェクト生成

クラスのコンストラクタに依存せずにオブジェクトを作成できるため、実行時に動的にオブジェクトの複製が可能です。これにより、コードの柔軟性が向上します。

##### カプセル化の強化

オブジェクトの内部構造や具象クラスを意識せずにコピーできるため、実装の詳細を隠蔽しながらオブジェクトの生成が可能になります。

#### まとめ

Prototype パターンを使用することで、オブジェクト生成のコストを削減し、柔軟かつカプセル化された形でオブジェクトを作成できます。特に、初期化コストが高いオブジェクトや、動的なインスタンス生成が求められる場面で有効です。

### Builder

#### 概要

Builder パターンは、複雑なオブジェクトの生成を簡潔かつ柔軟に行うためのデザインパターンです。オブジェクトの構築過程をカプセル化し、ステップごとに構成できるようにします。

#### メリット

##### 柔軟なオブジェクト生成

オブジェクトの作成手順を分離することで、同じ構築プロセスから異なる構成のオブジェクトを生成できます。

##### 可読性の向上

メソッドチェーンを活用することで、オブジェクトの設定内容を直感的に記述できます。

##### 不変オブジェクトの構築

必要なフィールドを設定した後にオブジェクトを生成するため、イミュータブルなオブジェクトを容易に作成できます。

##### コンストラクタの複雑化を回避

コンストラクタの引数が多くなる問題を避け、オプションのパラメータを適切に管理できます。

#### まとめ

Builder パターンを使用することで、複雑なオブジェクトの生成を柔軟かつ分かりやすく行うことができます。特に、多くのオプションを持つオブジェクトの作成時に有用です。

### Abstract Factory

#### 概要

Abstract Factory パターンは、関連する一連のオブジェクトを生成するためのインターフェースを提供し、具体的な生成クラスを指定せずに、オブジェクトのファミリーを作成できるデザインパターンです。

#### メリット

##### 疎結合の実現

- 具体的な実装クラスから独立したコードを作成できるため、システムの柔軟性が大幅に向上します。
- クライアントコードは具体的な生成クラスと密結合せず、抽象インターフェースを通じて操作できるため、コンポーネント間の依存関係を最小限に抑えられます。

##### 一貫性のあるオブジェクト生成

- 関連するオブジェクトのセットを常に一緒に生成することができ、オブジェクト間の整合性を保証します。
- 異なるファミリーのオブジェクトを簡単に切り替えることができ、プラットフォームや環境に応じた柔軟な対応が可能になります。

##### コードの拡張性と保守性の向上

- 新しい製品ファミリーを追加する際に、既存のコードを変更することなく拡張できます。
- 具体的な実装の詳細をクライアントコードから隠蔽することで、システムの変更や修正が容易になります。

##### テスタビリティの向上

- モックオブジェクトや代替実装を簡単に導入できるため、単体テストや統合テストが容易になります。
- 依存性の注入とテスト可能なコードデザインを促進します。

#### まとめ

Abstract Factory パターンは、複雑なシステムにおけるオブジェクト生成の柔軟性、拡張性、一貫性を実現する強力なデザインパターンです。特に、異なるプラットフォームやバリエーションが存在するソフトウェアアーキテクチャにおいて、その威力を発揮します。

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

### Strategy

#### 概要

Strategy パターンは、アルゴリズムをカプセル化し、動的に切り替え可能にするデザインパターンです。特定の処理を抽象化し、異なる実装を持つ戦略（Strategy）として定義することで、コードの柔軟性と拡張性を向上させます。

#### メリット

##### アルゴリズムの分離とカプセル化

異なるアルゴリズムを個別のクラスとして分離できるため、メインのコードがシンプルになり、可読性や保守性が向上します。

##### 動的なアルゴリズムの切り替え

実行時に適切な戦略を選択できるため、条件分岐を減らし、柔軟な設計が可能になります。

##### Open/Closed Principle（開閉原則）の遵守

新しい戦略を追加する際に既存のコードを変更せずに拡張できるため、変更に強い設計が可能です。

##### 単体テストが容易になる

アルゴリズムが独立したクラスとして実装されるため、個別にテストでき、テストの網羅性が向上します。

#### まとめ

Strategy パターンを使用することで、アルゴリズムの分離・動的切り替えが可能になり、コードの拡張性と保守性が向上します。また、新しい戦略を追加する際にも既存コードを変更せずに対応できるため、変更に強い設計を実現できます。

### Visitor

#### 概要

Visitor パターンは、オブジェクト構造内の各要素に対して新しい操作を追加できるようにする振る舞いに関するデザインパターンです。このパターンは、データ構造とアルゴリズムを分離し、オブジェクトの構造を変更することなく新しい操作を追加できる柔軟な方法を提供します。

#### メリット

##### 開放閉鎖の原則の遵守

- 既存のオブジェクト構造を変更せずに、新しい操作を簡単に追加できます。
- クラス階層を修正することなく、新しい振る舞いを導入できるため、ソフトウェアの拡張性が向上します。

##### 関心の分離

- 各オブジェクトの構造と、それらに対する操作を明確に分離することができます。
- 異なる操作を個別の Visitor クラスに encapsulate できるため、コードの整理と保守性が向上します。

##### 複雑な操作の簡素化

- 複数の異なる型のオブジェクトに対して、一貫した方法で操作を実行できます。
- ダブルディスパッチを利用することで、型に応じた柔軟な処理が可能になります。

##### コードの再利用性

- 共通の振る舞いを持つ操作を、簡単に再利用および共有できます。
- 新しい操作の追加が容易になり、コードの重複を最小限に抑えられます。

#### まとめ

Visitor パターンは、オブジェクト構造と操作を分離することで、柔軟で拡張性の高いコード設計を可能にします。特に、頻繁に新しい操作を追加する必要がある複雑なオブジェクト構造を扱う際に有効なパターンです。ただし、過度に使用すると複雑さが増す可能性があるため、適切なユースケースを見極めることが重要です。

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

### Bridge

#### 概要

Bridge パターンは、抽象化と実装を分離し、両者を独立して変更できるようにするデザインパターンです。オブジェクトの抽象部分と実装部分を別々の階層に分けることで、柔軟性とモジュール性を高めることができます。

#### メリット

##### 柔軟な設計と拡張性

- 抽象化と実装を分離することで、各レイヤーを独立して拡張できます。
- 新しい抽象クラスや実装クラスを追加する際に、既存のコードを大幅に修正する必要がありません。
- 実行時に実装を動的に切り替えることが可能になります。

##### 関心の分離

- 抽象化（インターフェース）と実装（具象クラス）を明確に分離することで、各コンポーネントの責任範囲が明確になります。
- コードの可読性と保守性が向上し、開発者が個々のコンポーネントに集中できます。

##### 継承の代替手段

- 多重継承による複雑さを避けつつ、柔軟な機能拡張が可能です。
- クラス爆発（combinatorial explosion）の問題を解決し、複雑な継承階層を簡素化できます。

##### プラットフォームや実装の独立性

- 異なるプラットフォームや実装間の依存関係を低減できます。
- 新しいプラットフォームや実装を追加する際に、既存のコードへの影響を最小限に抑えられます。

#### まとめ

Bridge パターンは、オブジェクトの抽象化と実装を分離することで、システムの柔軟性、拡張性、保守性を大幅に向上させます。複雑な継承構造を避けながら、コンポーネント間の結合を緩和し、より modularity の高いソフトウェア設計を実現できます。

### Composite

#### 概要

Composite パターンは、オブジェクトをツリー構造で管理し、個々のオブジェクトとグループ化されたオブジェクトを同一視できるデザインパターンです。  
このパターンを使用することで、クライアントコードは個々の要素と複合要素を統一的に扱うことができます。

#### メリット

##### 階層構造をシンプルに表現できる

オブジェクトの階層構造を直感的に表現でき、親子関係を適切に管理できるようになります。

##### クライアントコードの単純化

個々の要素と複合要素を統一的に扱えるため、クライアントコードが複雑な分岐処理を持たずに済みます。

##### 拡張性が高い

新しい要素を追加する際に、既存のコードを大きく変更する必要がなく、柔軟に拡張できます。

#### まとめ

Composite パターンを利用することで、オブジェクトの階層構造を適切に管理し、コードの可読性や拡張性を向上させることができます。

### Decorator

#### 概要

Decorator パターンは、既存のオブジェクトの振る舞いを変更・拡張するための設計パターンです。継承を使用せずに動的に機能を追加できるため、柔軟な設計が可能になります。

#### メリット

##### クラスの継承を増やさずに機能を拡張できる

Decorator パターンを使うことで、機能を追加するために新しいサブクラスを作成する必要がなくなります。継承の階層が深くならず、可読性や保守性が向上します。

##### 柔軟に機能を組み合わせられる

複数のデコレーターを組み合わせることで、必要な機能だけを選択的に適用できます。これにより、クラスの組み合わせ爆発を防ぎつつ、機能のカスタマイズが可能になります。

##### Open/Closed Principle（OCP）を遵守できる

既存のクラスを修正せずに新しい機能を追加できるため、OCP（拡張には開かれているが、修正には閉じている）の原則に従った設計ができます。

##### 実行時に動的な変更が可能

Decorator を使うことで、実行時にオブジェクトへ追加機能を付与できます。これにより、状況に応じた振る舞いの変更が容易になります。

#### まとめ

Decorator パターンを使うことで、継承に頼らずに柔軟かつ動的にオブジェクトの振る舞いを拡張できます。OCP を守りながら、機能の組み合わせを容易にし、保守性の高いコードを実現できます。

## README ChatGPT/Claude 出力プロンプト

```Markdown
xxxパターンを使用することによって何が嬉しいのかを教えてください。
READMEに追記したいので、以下の形式を使用してMarkdown形式で書いてください。

### xxx
#### 概要
#### メリット
##### {具体的なメリットの見出し}
#### まとめ
```
