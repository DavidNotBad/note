<?php

date_default_timezone_set('prc');


//echo 'now:', MyDateTime::now()->toDateTimeString(), '<br>';
//echo 'createTomorrow:', MyDateTime::createTomorrow(), '<br>';
//echo 'createYesterday:', MyDateTime::createYesterday(), '<br>';
//echo 'createThisSunday:', MyDateTime::createThisSunday(), '<br>';
//echo 'createNextSunday:', MyDateTime::createNextSunday(), '<br>';
//echo 'createLastSunday:', MyDateTime::createLastSunday(), '<br>';
//echo 'getSecond:', MyDateTime::now()->getSecond(), '<br>';
//echo 'setSecond:', MyDateTime::now()->setSecond(10), '<br>';
//echo 'startOfDay:', MyDateTime::now()->startOfDay(), '<br>';
//echo 'endOfDay:', MyDateTime::now()->endOfDay(), '<br>';
//echo 'nextMonday+lastMonday:', MyDateTime::now()->nextMonday()->lastMonday(), '<br>';
//echo 'startOfWeek:', MyDateTime::now()->startOfWeek(), '<br>';
//echo 'endOfWeek:', MyDateTime::now()->endOfWeek(), '<br>';
//echo 'startOfMonth:', MyDateTime::now()->startOfMonth(), '<br>';
//echo 'endOfMonth:', MyDateTime::now()->endOfMonth(), '<br>';
//echo 'getWeekOfMonth:', MyDateTime::now()->getWeekOfMonth(), '<br>';
//echo 'startOfYear:', MyDateTime::now()->startOfYear(), '<br>';
//echo 'endOfYear:', MyDateTime::now()->endOfYear(), '<br>';
//echo 'lastMonth:', MyDateTime::now()->lastMonth(), '<br>';
//echo 'getWeekOfMonth:', MyDateTime::now()->getWeekOfMonth(), '<br>';



/**
 * Class MyDateTime
 */
class MyDateTime{
    /**
     * 格式化输出的指定格式
     */
    const PRC = 'Y-m-d H:i:s';

    /**
     * 一分钟所包含的秒数
     */
    const MINUTE = 60;
    /**
     * 一小时所包含的秒数
     */
    const HOUR = 3600;
    /**
     * 一天所包含的秒数
     */
    const DAY = 86400;

    /**
     * 周日
     */
    const SUNDAY = 0;
    /**
     * 周一
     */
    const MONDAY = 1;
    /**
     * 周二
     */
    const TUESDAY = 2;
    /**
     * 周三
     */
    const WEDNESDAY = 3;
    /**
     * 周四
     */
    const THURSDAY = 4;
    /**
     * 周五
     */
    const FRIDAY = 5;
    /**
     * 周六
     */
    const SATURDAY = 6;


    /**
     * @var false|int   时间戳
     */
    private $timestamp;
    /**
     * @var bool    一周是否从星期天开始计算
     */
    private static $isWeekBeginsOnSunday = false;

    /**
     * MyDateTime constructor.
     * @param $timestamp
     */
    public function __construct($timestamp) {
        if($timestamp instanceof $this){
            $timestamp = $timestamp->getTimestamp();
        }elseif(!($timestamp && $this->isInterger($timestamp) && $timestamp > 0)){
            $timestamp = strtotime($timestamp);
        }
        $this->timestamp = $timestamp;
    }

    /**
     * 用指定的时间创建一个实例
     *
     * @param $year
     * @param $month
     * @param $day
     * @param $hour
     * @param $minute
     * @param $second
     * @return MyDateTime
     */
    public static function make($year, $month, $day, $hour, $minute, $second)
    {
        $time = mktime($hour, $minute, $second, $month, $day, $year);
        return static::createFromTimestamp($time);
    }

    /**
     * 用时间戳创建一个实例
     *
     * @param $timestamp
     * @return static
     */
    public static function createFromTimestamp($timestamp)
    {
        return new static($timestamp);
    }

    /**
     * 获取包含当前时间的实例
     *
     * @return MyDateTime
     */
    public static function now()
    {
        $time = time();
        return static::createFromTimestamp($time);
    }

    /**
     * 获取明天的实例
     *
     * @return MyDateTime
     */
    public static function createTomorrow()
    {
        return static::now()->addDays(1);
    }

    /**
     * 获取昨天的实例
     *
     * @return MyDateTime
     */
    public static function createYesterday()
    {
        return static::now()->subDays(1);
    }

    /**
     * 获取本周日的实例
     *
     * @return MyDateTime
     */
    public static function createThisSunday()
    {
        return static::now()->addSundays(0);
    }

    /**
     * 获取下周日的实例
     *
     * @return MyDateTime
     */
    public static function createNextSunday()
    {
        return static::now()->addSundays(1);
    }

    /**
     * 获取上周日的实例
     *
     * @return MyDateTime
     */
    public static function createLastSunday()
    {
        return static::now()->subSundays(1);
    }

    /**
     * 获取本周一的实例
     *
     * @return MyDateTime
     */
    public static function createThisMonday()
    {
        return static::now()->addMondays(0);
    }

    /**
     * 获取下周一的实例
     *
     * @return MyDateTime
     */
    public static function createNextMonday()
    {
        return static::now()->addMondays(1);
    }

    /**
     * 获取上周一的实例
     *
     * @return MyDateTime
     */
    public static function createLastMonday()
    {
        return static::now()->subMondays(1);
    }

    /**
     * 获取本周二的实例
     *
     * @return MyDateTime
     */
    public static function createThisTuesday()
    {
        return static::now()->addTuesdays(0);
    }

    /**
     * 获取下周二的实例
     *
     * @return MyDateTime
     */
    public static function createNextTuesday()
    {
        return static::now()->addTuesdays(1);
    }

    /**
     * 获取上周二的实例
     *
     * @return MyDateTime
     */
    public static function createLastTuesday()
    {
        return static::now()->subTuesdays(1);
    }

    /**
     * 获取本周三的实例
     *
     * @return MyDateTime
     */
    public static function createThisWednesday()
    {
        return static::now()->addWednesdays(0);
    }

    /**
     * 获取下周三的实例
     *
     * @return MyDateTime
     */
    public static function createNextWednesday()
    {
        return static::now()->addWednesdays(1);
    }

    /**
     * 获取上周三的实例
     *
     * @return MyDateTime
     */
    public static function createLastWednesday()
    {
        return static::now()->subWednesdays(1);
    }

    /**
     * 获取本周四的实例
     *
     * @return MyDateTime
     */
    public static function createThisThursday()
    {
        return static::now()->addThursdays(0);
    }

    /**
     * 获取下周四的实例
     *
     * @return MyDateTime
     */
    public static function createNextThursday()
    {
        return static::now()->addThursdays(1);
    }

    /**
     * 获取上周四的实例
     *
     * @return MyDateTime
     */
    public static function createLastThursday()
    {
        return static::now()->subThursdays(1);
    }

    /**
     * 获取这周五的实例
     *
     * @return MyDateTime
     */
    public static function createThisFriday()
    {
        return static::now()->addFridays(0);
    }

    /**
     * 获取下周五的实例
     *
     * @return MyDateTime
     */
    public static function createNextFriday()
    {
        return static::now()->addFridays(1);
    }

    /**
     * 获取上周五的实例
     *
     * @return MyDateTime
     */
    public static function createLastFriday()
    {
        return static::now()->subFridays(1);
    }

    /**
     * 获取本周六的实例
     *
     * @return MyDateTime
     */
    public static function createThisSaturday()
    {
        return static::now()->addSaturdays(0);
    }

    /**
     * 获取下周六的实例
     *
     * @return MyDateTime
     */
    public static function createNextSaturday()
    {
        return static::now()->addSaturdays(1);
    }

    /**
     * 获取上周六的实例
     *
     * @return MyDateTime
     */
    public static function createLastSaturday()
    {
        return static::now()->subSaturdays(1);
    }

    /**
     * 设置一周是否开始于星期天
     *
     * @param $bool
     */
    public static function isWeekBeginsOnSunday($bool)
    {
        static::$isWeekBeginsOnSunday = (bool) $bool;
    }


    /**
     * 把时间设置为明天
     *
     * @return MyDateTime
     */
    public function tomorrow()
    {
        return $this->addDays(1);
    }

    /**
     * 把时间设置成昨天
     *
     * @return MyDateTime
     */
    public function yesterday()
    {
        return $this->subDays(1);
    }

    /**
     * 把时间设置成本周日
     *
     * @return MyDateTime
     */
    public function thisSunday()
    {
        return $this->addSundays(0);
    }

    /**
     * 把时间设置成下周日
     *
     * @return MyDateTime
     */
    public function nextSunday()
    {
        return $this->addSundays(1);
    }

    /**
     * 把时间设置成上周日
     *
     * @return MyDateTime
     */
    public function lastSunday()
    {
        return $this->subSundays(1);
    }

    /**
     * 把时间设置成本周一
     *
     * @return MyDateTime
     */
    public function thisMonday()
    {
        return $this->addMondays(0);
    }

    /**
     * 把时间设置成下周一
     *
     * @return MyDateTime
     */
    public function nextMonday()
    {
        return $this->addMondays(1);
    }

    /**
     * 把时间设置成上周一
     *
     * @return MyDateTime
     */
    public function lastMonday()
    {
        return $this->subMondays(1);
    }

    /**
     * 把时间设置成本周二
     *
     * @return MyDateTime
     */
    public function thisTuesday()
    {
        return $this->addTuesdays(0);
    }

    /**
     * 把时间设置成下周二
     *
     * @return MyDateTime
     */
    public function nextTuesday()
    {
        return $this->addTuesdays(1);
    }

    /**
     * 把时间设置成上周二
     *
     * @return MyDateTime
     */
    public function lastTuesday()
    {
        return $this->subTuesdays(1);
    }

    /**
     * 把时间设置成本周三
     *
     * @return MyDateTime
     */
    public function thisWednesday()
    {
        return $this->addWednesdays(0);
    }

    /**
     * 把时间设置成下周三
     *
     * @return MyDateTime
     */
    public function nextWednesday()
    {
        return $this->addWednesdays(1);
    }

    /**
     * 把时间设置成上周三
     *
     * @return MyDateTime
     */
    public function lastWednesday()
    {
        return $this->subWednesdays(1);
    }

    /**
     * 把时间设置成本周四
     *
     * @return MyDateTime
     */
    public function thisThursday()
    {
        return $this->addThursdays(0);
    }

    /**
     * 把时间设置成下周四
     *
     * @return MyDateTime
     */
    public function nextThursday()
    {
        return $this->addThursdays(1);
    }

    /**
     * 把时间设置成上周四
     *
     * @return MyDateTime
     */
    public function lastThursday()
    {
        return $this->subThursdays(1);
    }

    /**
     * 把时间设置成本周五
     *
     * @return MyDateTime
     */
    public function thisFriday()
    {
        return $this->addFridays(0);
    }

    /**
     * 把时间设置成下周五
     *
     * @return MyDateTime
     */
    public function nextFriday()
    {
        return $this->addFridays(1);
    }

    /**
     * 把时间设置成上周五
     *
     * @return MyDateTime
     */
    public function lastFriday()
    {
        return $this->subFridays(1);
    }

    /**
     * 把时间设置成本周六
     *
     * @return MyDateTime
     */
    public function thisSaturday()
    {
        return $this->addSaturdays(0);
    }

    /**
     * 把时间设置成下周六
     *
     * @return MyDateTime
     */
    public function nextSaturday()
    {
        return $this->addSaturdays(1);
    }

    /**
     * 把时间设置成上周六
     *
     * @return MyDateTime
     */
    public function lastSaturday()
    {
        return $this->subSaturdays(1);
    }

    /**
     * 设置当前秒数
     *
     * @param $second
     * @return MyDateTime
     */
    public function setSecond($second)
    {
        $thisSecond = $this->getSecond();
        return $this->addSeconds($second - $thisSecond);
    }

    /**
     * 获取当前秒数
     *
     * @return false|string
     */
    public function getSecond()
    {
        return date('s', $this->timestamp);
    }

    /**
     * 设置当前分钟数
     *
     * @param $minute
     * @return MyDateTime
     */
    public function setMinute($minute)
    {
        $thisMinute = $this->getMinute();
        return $this->addMinutes($minute - $thisMinute);
    }

    /**
     * 获取当前分钟数
     *
     * @return false|string
     */
    public function getMinute()
    {
        return date('i', $this->timestamp);
    }

    /**
     * 设置当前小时数
     *
     * @param $hour
     * @return MyDateTime
     */
    public function setHour($hour)
    {
        $thisHour = $this->getHour();
        return $this->addHours($hour - $thisHour);
    }

    /**
     * 获取当前小时数
     *
     * @return false|string
     */
    public function getHour()
    {
        return date('H', $this->timestamp);
    }

    /**
     * 设置当前天数
     *
     * @param $day
     * @return MyDateTime
     */
    public function setDay($day)
    {
        $thisDay = $this->getDay();
        return $this->addDays($day - $thisDay);
    }

    /**
     * 获取当前天数
     *
     * @return false|string
     */
    public function getDay()
    {
        return date('d', $this->timestamp);
    }

    /**
     * 设置当前月数
     *
     * @param $month
     * @return $this
     */
    public function setMonth($month)
    {
        $this->timestamp = $this->mktime(
            $this->getYear(),
            $month,
            $this->getDay(),
            $this->getHour(),
            $this->getMinute(),
            $this->getSecond()
        );
        return $this;
    }

    /**
     * 获取当前月数
     *
     * @return false|string
     */
    public function getMonth()
    {
        return date('m', $this->timestamp);
    }

    /**
     * 设置当前年数
     *
     * @param $year
     * @return $this
     */
    public function setYear($year)
    {
        $this->timestamp = $this->mktime(
            $year,
            $this->getMonth(),
            $this->getDay(),
            $this->getHour(),
            $this->getMinute(),
            $this->getSecond()
        );
        return $this;
    }

    /**
     * 获取当前年数
     *
     * @return false|string
     */
    public function getYear()
    {
        return date('Y', $this->timestamp);
    }

    /**
     * 使用时间戳重置当前时间
     *
     * @param $timestamp
     * @return $this
     */
    public function setByTimestamp($timestamp)
    {
        $this->timestamp = $timestamp;
        return $this;
    }

    /**
     * 使用字符串重置当前时间
     *
     * @param $timeString
     * @return $this
     */
    public function setByString($timeString)
    {
        $timestamp = strtotime($timeString);
        $this->timestamp = $timestamp;
        return $this;
    }

    /**
     * 把当前时间设置成一天的开始
     *
     * @return MyDateTime
     */
    public function startOfDay()
    {
        return $this->setHour(0)->setMinute(0)->setSecond(0);
    }

    /**
     * 把当前时间设置成一天的结束
     *
     * @return MyDateTime
     */
    public function endOfDay()
    {
        return $this->setHour(23)->setMinute(59)->setSecond(59);
    }


    /**
     * 把当前时间设置成一周的开始
     *
     * @return MyDateTime
     */
    public function startOfWeek()
    {
        if(static::$isWeekBeginsOnSunday) {
            return $this->thisSunday()->startOfDay();
        }else{
            return $this->thisMonday()->startOfDay();
        }
    }

    /**
     * 把当前时间设置成一周的结束
     *
     * @return MyDateTime
     */
    public function endOfWeek()
    {
        if(static::$isWeekBeginsOnSunday) {
            return $this->thisSaturday()->endOfDay();
        }else{
            return $this->thisSunday()->endOfDay();
        }
    }

    /**
     * 把当前时间设置成一个月的开始
     *
     * @return MyDateTime
     */
    public function startOfMonth()
    {
        return $this->setDay(1)->startOfDay();
    }

    /**
     * 获取当前月有几个星期
     *
     * @return int
     */
    public function getWeekOfMonth()
    {
        //第一周的最后一天是本月的第几天
        $nthMonthDays = $this->startOfMonth()->endOfWeek()->getNthMonthDays();
        //一个月有几天
        $monthDays = $this->getMonthDays();
        return (int) ceil(($monthDays - $nthMonthDays) / 7) + 1;
    }

    /**
     * 把当前时间设置成一个月的结束
     *
     * @return MyDateTime
     */
    public function endOfMonth()
    {
        return $this->nextMonth()->startOfMonth()->subSeconds(1);
    }

    /**
     * 把当前时间设置成一年的开始
     *
     * @return MyDateTime
     */
    public function startOfYear()
    {
        return $this->setMonth(1)->startOfMonth();
    }

    /**
     * 把当前时间设置成一年的结束
     *
     * @return MyDateTime
     */
    public function endOfYear()
    {
        return $this->nextYear()->startOfYear()->subSeconds(1);
    }

    /**
     * 把当前时间设置为上一年
     *
     * @return MyDateTime
     */
    public function lastYear()
    {
        return $this->subYears(1);
    }

    /**
     * 把当前时间设置为下一年
     *
     * @return MyDateTime
     */
    public function nextYear()
    {
        return $this->addYears(1);
    }


    /**
     * 把当前时间设置成上个月
     *
     * @return MyDateTime
     */
    public function lastMonth()
    {
        return $this->subMonths(1);
    }

    /**
     * 把当前时间设置成下个月
     *
     * @return MyDateTime
     */
    public function nextMonth()
    {
        return $this->addMonths(1);
    }

    /**
     * 把当前时间设置成昨天
     *
     * @return MyDateTime
     */
    public function lastDay()
    {
        return $this->subDays(1);
    }

    /**
     * 把当前时间设置成明天
     *
     * @return MyDateTime
     */
    public function nextDay()
    {
        return $this->addDays(1);
    }

    /**
     * 把当前时间设置成一个小时前
     *
     * @return MyDateTime
     */
    public function lastHour()
    {
        return $this->subHours(1);
    }

    /**
     * 把当前时间设置成一个小时后
     *
     * @return MyDateTime
     */
    public function nextHour()
    {
        return $this->addHours(1);
    }

    /**
     * 把当前时间设置成一分钟前
     *
     * @return MyDateTime
     */
    public function lastMinute()
    {
        return $this->subMinutes(1);
    }

    /**
     * 把当前时间设置成一分钟后
     *
     * @return MyDateTime
     */
    public function nextMinute()
    {
        return $this->addMinutes(1);
    }

    /**
     * 把当前时间设置成上一秒
     *
     * @return MyDateTime
     */
    public function lastSecond()
    {
        return $this->subSeconds(1);
    }

    /**
     * 把当前时间设置成下一秒
     *
     * @return MyDateTime
     */
    public function nextSecond()
    {
        return $this->addSeconds(1);
    }


    /**
     * 魔术方法： 把类当成字符串使用时触发该函数
     *
     * @return false|string
     */
    public function __toString()
    {
        return $this->toDateTimeString();
    }

    /**
     * 格式化时间为日期格式
     *
     * @return false|string
     */
    public function toDateString()
    {
        return date('Y-m-d', $this->timestamp);
    }

    /**
     * 格式化时间为时间格式
     *
     * @return false|string
     */
    public function toTimeString()
    {
        return date('H:i:s', $this->timestamp);
    }

    /**
     * 格式化时间为日期时间格式
     *
     * @return false|string
     */
    public function toDateTimeString()
    {
        return date(static::PRC, $this->timestamp);
    }

    /**
     * 添加秒数
     *
     * @param $number
     * @return $this
     */
    public function addSeconds($number)
    {
        $this->timestamp += $number;
        return $this;
    }


    /**
     * 减少秒数
     *
     * @param $number
     * @return $this
     */
    public function subSeconds($number)
    {
        $this->timestamp -= $number;
        return $this;
    }

    /**
     * 添加分钟数
     *
     * @param $number
     * @return $this
     */
    public function addMinutes($number)
    {
        $this->timestamp += $number * static::MINUTE;
        return $this;
    }

    /**
     * 减少分钟数
     *
     * @param $number
     * @return $this
     */
    public function subMinutes($number)
    {
        $this->timestamp -= $number * static::MINUTE;
        return $this;
    }

    /**
     * 添加小时数
     *
     * @param $number
     * @return $this
     */
    public function addHours($number)
    {
        $this->timestamp += $number * static::HOUR;
        return $this;
    }

    /**
     * 减少小时数
     *
     * @param $number
     * @return $this
     */
    public function subHours($number)
    {
        $this->timestamp -= $number * static::HOUR;
        return $this;
    }

    /**
     * 添加天数
     *
     * @param $number
     * @return $this
     */
    public function addDays($number)
    {
        $this->timestamp += $number * static::DAY;
        return $this;
    }

    /**
     * 减少天数
     *
     * @param $number
     * @return $this
     */
    public function subDays($number)
    {
        $this->timestamp -= $number * static::DAY;
        return $this;
    }


    /**
     * 添加周数
     *
     * @param $number
     * @return MyDateTime
     */
    public function addWeeks($number)
    {
        return $this->addDays(7 * $number);
    }

    /**
     * 减少周数
     *
     * @param $number
     * @return MyDateTime
     */
    public function subWeeks($number)
    {
        return $this->addWeeks(-$number);
    }

    /**
     * 添加月份
     *
     * @param $number
     * @return MyDateTime
     */
    public function addMonths($number)
    {
        return $this->setMonth($this->getMonth() + $number);
    }

    /**
     * 减少月份
     *
     * @param $number
     * @return MyDateTime
     */
    public function subMonths($number)
    {
        return $this->addMonths(-$number);
    }


    /**
     * 添加年份
     *
     * @param $number
     * @return MyDateTime
     */
    public function addYears($number)
    {
        return $this->setYear($this->getYear() + $number);
    }

    /**
     * 减少年份
     *
     * @param $number
     * @return MyDateTime
     */
    public function subYears($number)
    {
        return $this->addYears(-$number);
    }

    /**
     * 同一周内，当前星期距离指定星期差几天
     *
     * @param $toDayNumber
     * @return false|int|string
     */
    public function offsetWeek($toDayNumber)
    {
        $weekNumber = $this->getWeekNumber();

        if(!static::$isWeekBeginsOnSunday) {
            if($weekNumber == 0){
                $weekNumber = 7;
            }
            if($toDayNumber == 0){
                $toDayNumber = 7;
            }
        }
        return $toDayNumber - $weekNumber;
    }

    /**
     * 增加指定数量个周日
     *
     * @param $number
     * @return MyDateTime
     */
    public function addSundays($number)
    {
        return $this->changeWeek(static::SUNDAY, $number);
    }

    /**
     * 减少指定数量个周日
     *
     * @param $number
     * @return MyDateTime
     */
    public function subSundays($number)
    {
        return $this->addSundays(-$number);
    }

    /**
     * 增加指定数量个周一
     *
     * @param $number
     * @return MyDateTime
     */
    public function addMondays($number)
    {
        return $this->changeWeek(static::MONDAY, $number);
    }

    /**
     * 减少指定数量个周一
     *
     * @param $number
     * @return MyDateTime
     */
    public function subMondays($number)
    {
        return $this->addMondays(-$number);
    }

    /**
     * 增加指定数量个周二
     *
     * @param $number
     * @return MyDateTime
     */
    public function addTuesdays($number)
    {
        return $this->changeWeek(static::TUESDAY, $number);
    }

    /**
     * 减少指定数量个周二
     *
     * @param $number
     * @return MyDateTime
     */
    public function subTuesdays($number)
    {
        return $this->addTuesdays(-$number);
    }

    /**
     * 增加指定数量个周三
     *
     * @param $number
     * @return MyDateTime
     */
    public function addWednesdays($number)
    {
        return $this->changeWeek(static::WEDNESDAY, $number);
    }

    /**
     * 减少指定数量个周三
     *
     * @param $number
     * @return MyDateTime
     */
    public function subWednesdays($number)
    {
        return $this->addWednesdays(-$number);
    }

    /**
     * 增加指定数量个周四
     *
     * @param $number
     * @return MyDateTime
     */
    public function addThursdays($number)
    {
        return $this->changeWeek(static::THURSDAY, $number);
    }

    /**
     * 减少指定数量个周四
     *
     * @param $number
     * @return MyDateTime
     */
    public function subThursdays($number)
    {
        return $this->addThursdays(-$number);
    }

    /**
     * 增加指定数量个周五
     *
     * @param $number
     * @return MyDateTime
     */
    public function addFridays($number)
    {
        return $this->changeWeek(static::THURSDAY, $number);
    }

    /**
     * 减少指定数量个周五
     *
     * @param $number
     * @return MyDateTime
     */
    public function subFridays($number)
    {
        return $this->addFridays(-$number);
    }

    /**
     * 增加指定数量个周六
     *
     * @param $number
     * @return MyDateTime
     */
    public function addSaturdays($number)
    {
        return $this->changeWeek(static::SATURDAY, $number);
    }

    /**
     * 减少指定数量个周六
     *
     * @param $number
     * @return MyDateTime
     */
    public function subSaturdays($number)
    {
        return $this->addSaturdays(-$number);
    }

    /**
     * 获取当前时间戳
     *
     * @return false|int
     */
    public function getTimestamp()
    {
        return $this->timestamp;
    }

    /**
     * 获取星期的数字形式表示
     *
     * @return false|string
     */
    public function getWeekNumber()
    {
        return date('w', $this->timestamp);
    }

    /**
     * 获取当前月有几天
     *
     * @return false|string
     */
    public function getMonthDays()
    {
        return date('t', $this->timestamp);
    }

    /**
     * 获取当前日是月份的第几天
     *
     * @return false|string
     */
    public function getNthMonthDays()
    {
        return date('j', $this->timestamp);
    }

    /**
     * 克隆本对象
     *
     * @return MyDateTime
     */
    public function copy()
    {
        return clone $this;
    }

    /**
     * 使用指定的时间创建时间戳
     *
     * @param $year
     * @param $month
     * @param $day
     * @param $hour
     * @param $minute
     * @param $second
     * @return false|int
     */
    private function mktime($year, $month, $day, $hour, $minute, $second)
    {
        return mktime($hour, $minute, $second, $month, $day, $year);
    }

    /**
     * 判断一个数是否整数
     *
     * @param $num
     * @return bool
     */
    private function isInterger($num)
    {
        return !(!is_numeric($num) || strpos($num,".")!==false);
    }

    /**
     * 当前时间偏移指定天数，在偏移指定周数
     *
     * @param $toDayNumber
     * @param $number
     * @return MyDateTime
     */
    private function changeWeek($toDayNumber, $number)
    {
        $offset = $this->offsetWeek($toDayNumber);
        return $this->addDays($offset)->addWeeks($number);
    }


}
