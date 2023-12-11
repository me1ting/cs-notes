# 基本概念

计算机科学中的`charset`术语可参考[RFC2278](https://datatracker.ietf.org/doc/html/rfc2278)，有三个重要的术语：

1. `字符集`(charset)定义为将字节序列转换为字符序列的方法，是后面两个术语的组合，如使用UTF-8编码的Unicode。
2. `编码字符集`(coded character sets) 是抽象字符集合和整数集合之间的映射关系，如US-ASCII,ISO8859-1,Unicode。
3. `字符编码方案`(character-encoding scheme) 是编码字符集和字节序列集合之间的映射关系，如UTF-8,UTF-16,EUC。

某些软件将`字符集`与`编码字符集`等同，是因为在单字节时代(US-ASCII,ISO-8859)，并没有字符编码方案的概念。到了多字节时代，为了兼容US-ASCII以及方便存储和传输等，字符编码方案被提出。

# 理解相关概念

下面从一些常见的名称来理解这些概念：

## ASCII

现代计算机发明于美国，最开始只需要将英文字符转换为二进制数据，`ASCII`因此诞生。ASCII包含了英文字母、数字、半角符号、控制符。

ASCII又称为`US-ASCII`，使用定长的1个字节记录抽象字符映射的整数，实际使用了字节的7位用于编码，包含128个抽象字符，最高位用作奇偶校验位。

存在`扩展ASCII`(EASCII)，使用了最高位来表示字符。

术语ASCII指的是US-ASCII，不包括扩展部分。

小结：**ASCII是一种编码字符集。**

## ISO-8859

其它采用字母文字的语言，如法语、德语，它们的文字也需要编码为二进制。由于它们的字符集合不大，因此单字节剩余的容量足以包含。

`ISO-8859`是一系列标准，定义了15个采用8位整数的字符集，低位兼容ASCII，高位用于其它字母文字。不同的子标准，相同的高位表示不同的字符，因此它们是互不兼容的。

其中最常见的是`ISO 8859-1`，又称作`latin1`，其包含了西欧语言的大部分字符，是一些平台的默认编码字符集。

小结：**ISO-8859是一系列编码字符集。**

## GB2312,GBK

计算机传入中国后，汉字也需要编码为二进制数据。以汉字为代表的象形文字其抽象字符集合远远超出单字节所能容纳的大小。`GB2312-80`是中国官方发布的编码字符集标准，收录了汉字、拉丁字母、希腊字母、俄文字母、日文假名、符号等7445个图形字符。

GB2312字符集分成94个区，每区有94个位，共8836个码位。采用两个字节编码，分别记录区号和位号，如："啊"字位于16区1位，则是16-1。

但在工程上必须兼容ASCII，普遍采用的是`EUC-CN`表示法：这是一个变长的编码方案，对ASCII字符采用原来的单字节；对于GB2312字符，区号和位号都加上0xA0(160)。这样一个小于128的字节还是表示ASCII字符，而大于128的则一定表示的是GB2312的字符。因为GB2312收录了ASCII中已有的字符，所以存在全角、半角的区分（为了满足印刷业需求）。

没有被收录的汉字还有许多，因此中国官方又发布了`GBK`（汉字内码扩展规范）。GBK扩展了GB2312所包含的字符，达到21886个字符。

GBK默认采用EUC-CN表示法，完全兼容采用EUC-CN表达法的GB2312。区的数量，每个区的位数都进行了扩展。但第二个字节不再一定大于128，只有通过第一个字节判断编码的是ASCII字符还是GBK字符。

小结：

- **GB2312标准是一种编码字符集**
- **EUC-CN是一种字符编码方案，相比UTF-8、UTF-16显得默默无闻**
- **通常所指的GB2312是EUC-CN表示法与GB2312编码字符集的组合**

## Unicode,UTF-8,UTF-16,UTF-32

当计算机传入世界各国家或地区后，它们都根据自己的语言定义了字符集。传统编码如GB2312,ISO8859-1,big5等，它们支持英语（兼容ASCII）+本地语言的环境，却无法支持多语言环境，因此Unicode标准诞生。Unicode对现有的大部分抽象字符集合进行了整理、编码，从而实现多语言环境。

>Unicode标准是由Unicode联盟制定的，后来由国际标准化组织(ISO)进行标准化为ISO/IEC 10646系列标准，它们拥有相同的字符集和编码方案。

Unicode的编码空间为U+0000到U+10FFFF,共有1,112,064个码位(code point)。共划分为17个平面，其中U+0000到U+FFFF称为`基本多语言平面`(BMP)，其余称为`辅助平面`。

Unicode标准定义了编码字符集 `the Universal Coded Character Set` (`UCS`)。

UCS最开始的版本是`UCS-2`(2-byte Universal Character Set)。采用两个字节表示码位，只能编码BMP对应的字符。兼容标准ASCII（高位为0）。

但随着Unicode的扩展，辅助平面的添加，两个字节已经不足以表示所有码位。Unicode编码字符集的下一个版本是`UCS-4`(4-byte Universal Character Set)，使用4个字节表示码位，兼容`UCS-2（高位两字节为0）。

采用2字节或4字节的UCS相比ASCII的1个字节浪费存储空间，而且影响传输效率。因此Unicode定义了the Unicode Transformation Format (`UTF`) encodings实现针对Unicode的可变长度的编码方案，来解决上述问题。

`UTF-8`使用1~4个字节对UCS编码，其中：

- ASCII字符采用1个字节表示，兼容ASCII
- 大部分汉字采用3个字节表示

UTF-8保证了一个字符的字节序列不会包含在另一个字符的字节序列中。

`UTF-16`是UCS-2的超集，定义它并非为了传输和存储，而是对UCS-2扩展以支持辅助平面的字符。使用2或4个字节对UCS编码，2个字节的基本单元称为码元。对于BMP上的字符，一个码元数值上与码位相等。对于辅助平面上的字符采用两个码元来表示，称为前导代理和后导代理，它们将码位进行编码，保证了都不与BMP中有效字符的码位冲突。

`UTF-32`功能上与UCS-4相同。

## UTF中的BE,LE与BOM

大端(BE)模式小端(LE)模式指的是的字节序，在计算机底层中指数据在内存中的字节序，大端指数据的高字节存放在内存的低地址中，数据的低字节存放在内存的高地址中，而小端模式相反。

这里的BE,LE指的是编码单元在存储和传输过程中的字节序。UTF-16与UTF-32都存在大小端区分：

| Name           | UTF-8  | UTF-16      | UTF-16BE   | UTF-16LE      | UTF-32      | UTF-32BE   | UTF-32LE      |
| -------------- | ------ | ----------- | ---------- | ------------- | ----------- | ---------- | ------------- |
| Code unit size | 8 bits | 16 bits     | 16 bits    | 16 bits       | 32 bits     | 32 bits    | 32 bits       |
| Byte order     | N/A    | &lt;BOM&gt; | big-endian | little-endian | &lt;BOM&gt; | big-endian | little-endian |

而UTF-8的编码单元为一个字节，因此不需要提供大小端模式。

UTF-16和UTF-32的编码单元分别为2个字节和4个字节，因此提供了3种风格。

`<BOM>`指由BOM决定字节序，如果未使用BOM，默认大端模式。

BOM(byte order mark)是一个Unicode字符，UTF-8,UTF-16,UTF-32都可以使用BOM。其作为魔数表明其所在文本：

- 采用的哪种字节序
- 该文本采用的编码字符集为Unicode
- 采用的哪种Unicode编码方案

> 大小端和BOM的使用场景不多，主要是Windows平台在使用。

# 结合具体环境理解

通过以上概念的学习总结后，回头看一些具体环境的charset与encoding，就很容易想明白了。

## Java(<=8)

`char`为2个字节大小，用来表示字符，因此Java的char只支持`UCS-2`中的字符。由于日常使用的汉字、英文、符号都位于基本面上，可以通过char表示，这种限制很少影响编程使用。但特殊情况，如许多Emoji字符将无法通过char来表示。

```java
char frowningFace='☹';//U+2639，正常编译
char slightlyFrowningFace='\uD83D\uDE41';//U+1F641 辅助平面字符🙁无法用char表示，无法编译
```

`String`表示字符串，使用UTF-16编码字符串并存储在底层的char\[]中。我们在编程时一般将底层char\[]的每个元素当作一个字符，进行比较或是匹配：

```java
if(str.charAt(0)=='<')
    ...
```

因为UTF-16的基本面字符码元与辅助字符码元不重复的性质，这种比较结果上并没有错误。只是注意得到的char不一定表示一个字符，而可能只是字符的部分。

理论上更好的方式，是以码位的方式进行遍历，但实际操作并不怎么方便：

```java
for(int i=0;i<faces.codePointCount(0,faces.length());i++){
    System.out.printf("%x\n",faces.codePointAt(i));
}
```

## Go

Go提供了字符类型`rune`，rune占4个字节，可以完整的表示所有的Unicode码位，因此rune完美的支持了`UCS-4`。

```
frowningFace,slightlyFrowningFace:='☹','🙁'//相比Java，Go的rune能表示所有Unicode字符
```

`string`为字符串类型，字符串通过`UTF-8`编码保存在底层的byte[]中。Go支持`for-range`对字符串中的Unicode字符进行迭代：

```
for pos, char := range str {
    ...
}
```

可以看到，作为较新的编程语言，Go对于字符编码的支持是足够现代化的。

## Windows(<=Win10 1607，新版本未测试)

这里只讨论Windows自带的文本编辑器`notepad`。notepad支持的编码方案包括：

- ASCII，统称所有的本地charset，如gb2312、big5等，与系统语言环境有关。
- Unicode，实际为带有BOM的UCS-2 LE。
- Unicode big endian，实际为带有BOM的UCS-2 BE。
- UTF-8，实际为带有BOM的UTF-8。

notepad对于编码的命名比较混乱，而且喜欢使用BOM，因此在Windows环境下最好不要使用notepad创建和编辑文本文档，而推荐使用其它软件，如Notepad++就很好的支持了不同的字符集和编码。

# 总结

理解charset的关键在于区分`charset`和`encoding`。对常用的US-ASCII,ISO8859-1,GB2312,GBK,UTF-8,UTF-16等有一个基本的概念，了解它们的历史发展和特点。

在使用US-ASCII、ISO8859-1、GB2312、GBK时我们不需要考虑编码方案的问题，因为它们的编码方式是唯一的，只有在Unicode时才需要考虑编码方案的问题。

不精确的讲，我们可以如下称呼：

- 字符集(charset)：US-ASCII,ISO8859-1,GB2312,GBK,Unicode
- 编码/编码方案(encoding)：EUC,UTF-8,UTF-16

# 参考资料

[java Charset class](https://docs.oracle.com/javase/7/docs/api/java/nio/charset/Charset.html)<br/>
[RFC2278](http://www.faqs.org/rfcs/rfc2278.html)<br/>
[wiki: ASCII](https://en.wikipedia.org/wiki/ASCII)<br/>
[wiki: ISO 8859](https://en.wikipedia.org/wiki/ISO%2FIEC_8859)<br/>
[wiki: GB2312](https://en.wikipedia.org/wiki/GB_2312)<br/>
[wiki: Extended Unix Code](https://en.wikipedia.org/wiki/Extended_Unix_Code)<br/>
[wiki: GBK](https://en.wikipedia.org/wiki/GBK)<br/>
[wiki: Unicode](https://en.wikipedia.org/wiki/Unicode)<br/>
[wiki: Universal_Coded_Character_Set](https://en.wikipedia.org/wiki/Universal_Coded_Character_Set)<br/>
[wiki: UTF_BOM](http://www.unicode.org/faq/utf_bom.html)<br/>
[全角和半角的区别？](https://www.zhihu.com/question/19605819)