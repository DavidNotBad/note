## 简单说明

```php
https://www.cnblogs.com/thors/p/9494057.html
```

## jTessBoxEditor训练字体

```php
https://www.cnblogs.com/cnlian/p/5765871.html
https://blog.csdn.net/zhou_zhu/article/details/78004131
https://blog.csdn.net/why200981317/article/details/48265621https://blog.csdn.net/why200981317/article/details/48265621
```

## php的tesseract-orc

```php
https://github.com/thiagoalessio/tesseract-ocr-for-php

[简单说明](https://www.cnblogs.com/thors/p/9494057.html)
[简单实用与训练](https://www.cnblogs.com/cnlian/p/5765871.html)
# tesserOCR 训练说明

git上已有开源的php类库实现了tesserOCR的方法  https://github.com/thiagoalessio/tesseract-ocr-for-php
Comporser 安装 comporser requir thiagoalessio/tesseract_ocr
其实只需要执行 exec 就可以。
<?php
new TesseractOCR('multi-languages.png')
->lang('eng', 'jpn', 'por') //使用语言包
->whitelist(range('A', 'Z')) //固定范围
 ->run();
```

## 中文文档

```php
https://qianjiye.de/2015/08/tesseract-ocr
```

