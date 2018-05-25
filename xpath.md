* 不包含某一个属性

    ```
    //tbody/tr[not(@class)]
    ```

* 后面元素

    ```
    //div/following-sibling::*[position()=1]
    ```

* 文本包含

  ```
  //div[contains(text(),'test4')]
  //div[contains(text(),'test3') or contains(text(),'test5')]
  ```


* 选中文本

  ```
  //div/text()
  ```

* 根据文本查找出当前元素在同类元素中的位置

  ```
  index-of (//div, //div[contains(text(),'test3')])
  ```

* 截取特定文本的位置

  ```
  从第一个到test4的位置
  subsequence(//div,1,index-of (//div, //div[contains(text(),'test4')]))

  从test2到test4的位置
  subsequence(//div,index-of (//div, //div[contains(text(),'test2')]),index-of (//div, //div[contains(text(),'test4')]) - index-of (//div, //div[contains(text(),'test2')]))
  ```

  ​

