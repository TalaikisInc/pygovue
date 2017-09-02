<template>
<div id="calendar" class="calendar">
  <div class="header">
    <i class="fa fa-fw fa-chevron-left" @click="subtractMonth"></i>
    <h4> </h4>
    <i class="fa fa-fw fa-chevron-right" @click="addMonth"></i>
  </div>
  <div>
    <ul>
      <li v-for="day in days" class="week"></li>
    </ul>
  </div>
  <div class="days">
  <ul>
    <li v-for="blank in firstDayOfMonth">&nbsp;</li>
    <li v-for="date in daysInMonth" class="day">{{ date }}</li>
  </ul>
  </div>
</div>
</template>

<script>
import moment from 'moment'

export default {
  data () {
    return {
      today: moment(),
      dateContext: moment(),
      days: ['S', 'M', 'T', 'W', 'T', 'F', 'S']
    }
  },
  methods: {
    addMonth () {
      var self = this
      self.dateContext = moment(self.dateContext).add(1, 'month')
    },
    subtractMonth () {
      var self = this
      self.dateContext = moment(self.dateContext).subtract(1, 'month')
    }
  },
  computed: {
    year () {
      var self = this
      return self.dateContext.format('Y')
    },
    month () {
      var self = this
      return self.dateContext.format('MMMM')
    },
    daysInMonth () {
      var self = this
      return self.dateContext.daysInMonth()
    },
    daysInWeek () {
      var self = this
      return self.dateContext.daysInWeek()
    },
    currentDate () {
      var self = this
      return self.dateContext.get('date')
    },
    firstDayOfMonth () {
      var self = this
      var firstDay = moment(self.dateContext).subtract((self.currentDate - 1), 'days')
      return firstDay.weekday()
    }
  }
}
</script>

<style scoped>
  .calendar {
    width: 350px;
    height: auto;
    overflow: hidden;
    -webkit-border-radius: 10px;
    border-radius: 10px;
    -webkit-box-shadow: 0 5px 10px 0 rgba(0, 0, 0, 0.2);
    box-shadow: 0 5px 10px 0 rgba(0, 0, 0, 0.2);
  }
  .calendar .header {
    width: 100%;
    padding: 10px 0;
    background-color: #2ecc71;
  }
  .calendar .head {
    text-align: center;
    font-size: 24px;
    padding: 10px 0;
    color: #ffffff;
    letter-spacing: 1px;
    text-shadow: 1px 1px 1px rgba(0, 0, 0, .1);
  }
  .calendar .weeks {
    display: block;
    width: 100%;
    overflow: auto;
    padding: 10px 0;
    text-align: center;
  }
  .calendar .weeks .week {
    width: 14.28571%;
    display: block;
    color: #ffffff;
    float: left;
    font-size: 16px;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
  .calendar .days {
    width: 100%;
    height: auto;
    overflow: auto;
    background-color: #ffffff;
    position: relative;
  }
  .calendar .days .day span{
    width: 50px;
    display: block;
    float: left;
    height: 50px;
    font-size: 12px;
    text-align: center;
    line-height: 50px;
    color: #333333;
    background-color: #fefefe;
    font-weight: bold;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
    border-right: 1px solid #f0f0f0;
    border-bottom: 1px solid #f0f0f0;
  }
  .calendar .days .day .this-month-day:hover{
    background-color: #e1e1e1;
    cursor: pointer;
    color: #ffffff;
  }
  .calendar .days .day .today {
    border-bottom: 3px solid #2ecc71;
    color: #2ecc71;
  }
  .calendar .days .day .not-this-month {
    background-color: #f9f9f9;
    color: #999999;
  }
  .calendar .input-group {
    background-color: #f9f9f9;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
  .calendar .input-group select {
  }
</style>