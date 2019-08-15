<?php

/**
 * Class IdcardGenerator
 * 身份证生成
 *
 * 参考标准
 * 正面:
 * “姓名”、“性别”、“民族”、“出生年月日”、“住址”、“公民身份号码”为6号黑体字，用蓝色油墨印刷；
 * 登记项目中的姓名项用5号黑体字（hei）印刷；其他项目则用小5号黑体字(hei)印刷；
 * 出生年月日 方正黑体简体(fzhei)字符大小：姓名＋号码（11点）其他（9点）字符间距（AV）：号码（50）字符行距：住址（12点）；
 * 身份证号码字体   OCR-B 10 BT(ocrb10bt)   文字 华文细黑。
 * 背面:
 * 左上角为国徽，用红色油墨印刷;
 * 其右侧为证件名称“中华人民共和国居民身份证”，分上下两排排列，其中上排的“中华人民共和国”为4号宋体字，下排的“居民身份证”为2号宋体字;
 * “签发机关”、“有效期限”为6号加粗黑体字;签发机关登记项采用，“xx市公安局”;
 * 有效期限采用“xxxx.xx-xxxx.xx.xx”格式，使用5号黑体字(hei)印刷，全部用黑色油墨印刷。
 */
class IdcardGenerator
{
    /**
     * @var string 姓名（5号黑体字）
     */
    protected $name;
    /**
     * @var string 生日-年（小5号黑体字）
     */
    protected $birthYear;
    /**
     * @var string 生日-月（小5号黑体字）
     */
    protected $birthMonth;
    /**
     * @var string 生日-日（小5号黑体字）
     */
    protected $birthDate;
    /**
     * @var string 民族（小5号黑体字）
     */
    protected $nation;
    /**
     * @var string 性别（小5号黑体字）
     */
    protected $gender;
    /**
     * @var string 地址（小5号黑体字）
     */
    protected $address;
    /**
     * @var string 身份证号码（OCR-B 10 BT字体）
     */
    protected $idCard;
    /**
     * @var string 头像地址
     */
    protected $headPortrait;
    /**
     * @var string 签发机关（5号黑体字）
     */
    protected $signingAndIssuingOrganization;
    /**
     * @var string 身份证有效期（5号黑体字）
     */
    protected $termOfValidity;
    /**
     * @var string 背景图片
     */
    protected $dstPath;
    /**
     * @var string 黑体字
     */
    protected $hei;
    /**
     * @var string ocr字体
     */
    protected $ocrb10bt;
    /**
     * @var resource gd库对象
     */
    protected $resource;
    /**
     * @var int 颜色
     */
    protected $color;

    public function __construct($dstPath, $hei, $ocrb10bt)
    {
        $this->setDstPath($dstPath);
        $this->setHei($hei);
        $this->setOcrb10bt($ocrb10bt);
        // 创建图片的实例
        $this->resource = imagecreatefromstring(file_get_contents($dstPath));
        //字体颜色
        $this->color = imagecolorallocate($this->resource, 10,10,10); // 字体颜色
    }

    /**
     * 获取正面照
     * @param $width
     * @param $height
     * @param null $filename
     * @return bool
     */
    public function getPositive($width, $height, $filename=null)
    {
        //填充信息
        $this->makePositive();

        //缩放裁剪图像
        $copyRes = $this->zip($width, $height, 260, 470, 1930, 1220);

        //自动保存或展示
        return $this->autoDown($copyRes, image_type_to_mime_type(IMAGETYPE_PNG), $filename);
    }


    public function getNegative($width, $height, $filename=null)
    {
        //填充信息
        $this->makeNegative();

        //缩放裁剪图像
        $copyRes = $this->zip($width, $height, 260, 470 + 1420, 1930, 1220);

        //自动保存或展示
        return $this->autoDown($copyRes, image_type_to_mime_type(IMAGETYPE_PNG), $filename);
    }

    protected function makeNegative()
    {
        //签发机关
        imagefttext($this->resource, 50, 0, 1075, 2810, $this->color, $this->hei, $this->signingAndIssuingOrganization);

        //有效期限(优化黑体字的“.”符号间距太大的问题)
        $termCount = strlen($this->termOfValidity);
        $offset = 0;
        for($i=0; $i<$termCount; $i++){
            $termTmp = substr($this->termOfValidity, $i, 1);
            imagefttext($this->resource, 50, 0, 1075 + 40 * $i - 20 * $offset, 2810 + 146, $this->color, $this->hei, $termTmp);
            if($termTmp == '.'){
                $offset++;
            }
        }
    }



    protected function zip($width, $height, $sx, $sy, $sw, $sh)
    {
        $copyRes = imagecreatetruecolor($width, $height);
        if($width >= $height){
            $dh = $height;
            $dw = $height * (240/151);
            $dx = ($width - $dw) / 2;
            $dy = 0;
        }else{
            $dw = $width;
            $dh = $width / (240/151);
            $dx = 0;
            $dy = ($height - $dh) / 2;
        }
        imagecopyresampled($copyRes, $this->resource, $dx, $dy, $sx, $sy, $dw, $dh, $sw, $sh);
        return $copyRes;
    }

    /**
     * 正面照补充信息
     */
    protected function makePositive()
    {
        //头像(把png图片伸缩到指定大小， 然后拷贝到身份证里)
        $tximg = $this->autoCreate($this->headPortrait);
        $tximgW = imagesx($tximg);
        $tximgH = imagesy($tximg);
        $newWidth = 500;
        $newHeight = 671;
        $newImg = imagecreatetruecolor($newWidth, $newHeight);
        $alpha = imagecolorallocatealpha($newImg, 0, 0, 0, 127);
        imagefill($newImg, 0, 0, $alpha);
        imagecopyresampled($newImg, $tximg, 0, 0, 0, 0, $newWidth, $newHeight, $tximgW, $tximgH);
        imagesavealpha($newImg, true);
        imagecopy($this->resource, $newImg, 1530, 700, 0, 0, $tximgW, $tximgH);
        imagedestroy($tximg);
        //姓名
        imagefttext($this->resource, 54, 0, 630, 760, $this->color, $this->hei, $this->name);
        //性别
        imagefttext($this->resource, 50, 0, 630, 760 + 140, $this->color, $this->hei, $this->gender);
        //民族
        imagefttext($this->resource, 50, 0, 630 + 400, 760 + 140, $this->color, $this->hei, $this->nation);
        //年
        imagefttext($this->resource, 50, 0, 630 + 20, 760 + 140 + 140 - 4, $this->color, $this->hei, $this->birthYear);
        //月
        imagefttext($this->resource, 50, 0, 630 + 20 + 290, 760 + 140 + 140 - 4, $this->color, $this->hei, $this->birthMonth);
        //日
        imagefttext($this->resource, 50, 0, 630 + 20 + 290 + 195, 760 + 140 + 140 - 4, $this->color, $this->hei, $this->birthDate);

        //住址
        $addressLen = mb_strlen($this->address, 'UTF-8');
        $addressSubCount = (int) ceil($addressLen / 11);
        $addressStart = 0;
        for ($i=0; $i<$addressSubCount; $i++){
            $addressTmp = mb_substr($this->address, $addressStart, 11, 'UTF-8');
            $addressStart += 11;
            imagefttext($this->resource, 50, 0, 630, 760 + 140 + 270 + 90 * $i, $this->color, $this->hei, $addressTmp);
        }

        //身份证号码
        $idCardCount = strlen($this->idCard);
        for($i=0; $i<$idCardCount; $i++){
            $idCardTmp = substr($this->idCard, $i, 1);
            //字体加粗
            imagefttext($this->resource, 50, 0, 630 + 345 + 50 * $i, 760 + 140 + 270 + 90 + 270, $this->color, $this->ocrb10bt, $idCardTmp);
            imagefttext($this->resource, 50, 0, 630 + 345 + 50 * $i + 1, 760 + 140 + 270 + 90 + 270 + 1, $this->color, $this->ocrb10bt, $idCardTmp);
        }

    }


    /**
     * @param mixed $dstPath
     */
    public function setDstPath($dstPath)
    {
        $this->dstPath = $dstPath;
    }

    /**
     * @param mixed $hei
     */
    public function setHei($hei)
    {
        $this->hei = $hei;
    }

    /**
     * @param mixed $ocrb10bt
     */
    public function setOcrb10bt($ocrb10bt)
    {
        $this->ocrb10bt = $ocrb10bt;
    }

    /**
     * @param string $name
     */
    public function setName($name)
    {
        $this->name = $name;
    }

    /**
     * @param string $birthYear
     */
    public function setBirthYear($birthYear)
    {
        $this->birthYear = $birthYear;
    }

    /**
     * @param string $birthMonth
     */
    public function setBirthMonth($birthMonth)
    {
        $birthMonth = str_pad($birthMonth, 2, ' ', STR_PAD_LEFT);
        $this->birthMonth = $birthMonth;
    }

    /**
     * @param string $birthDate
     */
    public function setBirthDate($birthDate)
    {
        $birthDate = str_pad($birthDate, 2, ' ', STR_PAD_LEFT);
        $this->birthDate = $birthDate;
    }

    /**
     * @param string $nation
     */
    public function setNation($nation)
    {
        $this->nation = $nation;
    }

    /**
     * @param string $gender
     */
    public function setGender($gender)
    {
        $this->gender = $gender;
    }

    /**
     * @param string $address
     */
    public function setAddress($address)
    {
        $this->address = $address;
    }

    /**
     * @param string $idCard
     */
    public function setIdCard($idCard)
    {
        $this->idCard = $idCard;
    }

    /**
     * @param string $headPortrait
     */
    public function setHeadPortrait($headPortrait)
    {
        $this->headPortrait = $headPortrait;
    }

    /**
     * @param string $signingAndIssuingOrganization
     */
    public function setSigningAndIssuingOrganization($signingAndIssuingOrganization)
    {
        $this->signingAndIssuingOrganization = $signingAndIssuingOrganization;
    }

    /**
     * @param string $termOfValidity
     */
    public function setTermOfValidity($termOfValidity)
    {
        $this->termOfValidity = $termOfValidity;
    }


    protected function autoCreate($path)
    {
        $imageSize = getimagesize($path);
        switch ($imageSize['mime']){
            case image_type_to_mime_type(IMAGETYPE_JPEG):
                return imagecreatefromjpeg($path);
                break;
            case image_type_to_mime_type(IMAGETYPE_PNG):
                return imagecreatefrompng($path);
                break;
            case image_type_to_mime_type(IMAGETYPE_GIF):
                return imagecreatefromgif($path);
                break;
            default:
                return false;
                break;
        }
    }

    protected function autoDown($image, $mine, $filename=null)
    {
        switch ($mine){
            case image_type_to_mime_type(IMAGETYPE_JPEG):
                header('Content-Type: image/jpeg');
                return imagejpeg($image, $filename);
                break;
            case image_type_to_mime_type(IMAGETYPE_PNG):
                header('Content-Type: image/png');
                return imagepng($image, $filename);
                break;
            case image_type_to_mime_type(IMAGETYPE_GIF):
                header('Content-Type: image/gif');
                return imagegif($image, $filename);
                break;
            default:
                return false;
                break;
        }
    }


}


header('Content-Type:text/html;charset=UTF-8');

$ig = new IdcardGenerator('./static/empty.png', './static/hei.ttf', './static/ocrb10bt.ttf');

$ig->setName('代用名');
$ig->setBirthYear('2013');
$ig->setBirthMonth('5');
$ig->setBirthDate('6');
$ig->setNation('汉');
$ig->setGender('男');
$ig->setAddress('湖南省长沙市开福区巡道街幸福小区居民组');
$ig->setIdCard('430512198908131367');
$ig->setHeadPortrait('./static/avatar.png');
$ig->setSigningAndIssuingOrganization('桃江县公安局');
$ig->setTermOfValidity('2013.03.05-2023.03.05');

$n = 2;
$ig->getPositive(240 * $n, 151 * $n, './a.png');
$ig->getNegative(240 * $n, 151 * $n, './b.png');