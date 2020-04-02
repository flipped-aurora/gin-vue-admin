<template>
  <div>
    <el-form :model="form" label-width="80px" ref="form">
      <el-row>
        <el-col  :span="3"><label for="">活动名称</label></el-col>
        <el-col  :span="10"><el-input v-model="form.name"></el-input></el-col>
      </el-row>
      <el-row>
        <el-col  :span="3"><label for="">活动区域</label></el-col>
        <el-col  :span="10">
           <el-select placeholder="请选择活动区域" v-model="form.region">
            <el-option label="上海" value="shanghai"></el-option>
            <el-option label="北京" value="beijing"></el-option>
           </el-select>
        </el-col>
      </el-row>
      <el-row>
        <el-col  :span="3"><label for="">活动时间</label></el-col>
        <el-col  :span="10">
           <el-col :span="11">
          <el-date-picker placeholder="选择日期" style="width: 100%;" type="date" v-model="form.date1"></el-date-picker>
        </el-col>
        <el-col :span="2" class="line">——</el-col>
        <el-col :span="11">
          <el-time-picker placeholder="选择时间" style="width: 100%;" v-model="form.date2"></el-time-picker>
        </el-col>
        </el-col>
      </el-row>
      <el-row>
        <el-col  :span="3"><label for="">即时配送</label></el-col>
        <el-col  :span="10"><el-switch v-model="form.delivery"></el-switch></el-col>
      </el-row>
      <el-row>
        <el-col  :span="3"><label for="">活动性质</label></el-col>
        <el-col  :span="10">
          <el-checkbox-group v-model="form.type">
            <el-checkbox label="美食/餐厅线上活动" name="type"></el-checkbox>
            <el-checkbox label="地推活动" name="type"></el-checkbox>
            <el-checkbox label="线下主题活动" name="type"></el-checkbox>
            <el-checkbox label="单纯品牌曝光" name="type"></el-checkbox>
          </el-checkbox-group>
        </el-col>
      </el-row>
      <el-row>
        <el-col  :span="3"><label for="">特殊资源</label></el-col>
        <el-col  :span="10">
          <el-radio-group v-model="form.resource">
            <el-radio label="线上品牌商赞助"></el-radio>
            <el-radio label="线下场地免费"></el-radio>
          </el-radio-group>
        </el-col>
      </el-row>
      <el-row>
        <el-col  :span="3"><label for="">活动形式</label></el-col>
        <el-col  :span="10"><el-input type="textarea" v-model="form.desc"></el-input></el-col>
      </el-row>
      <el-row>
        <el-col  :span="3"><label for="">穿梭框</label></el-col>
        <el-col  :span="20">
          <el-transfer
            :data="data"
            :filter-method="filterMethod"
            filter-placeholder="请输入城市拼音"
            filterable
            v-model="value"
          ></el-transfer>
        </el-col>
      </el-row>
      <el-row>
        <el-col  :span="3"><label for="">活动时间</label></el-col>
        <el-col  :span="10"><el-input v-model="form.name"></el-input></el-col>
      </el-row>
      <el-row type="flex" justify="center">
        <el-col :span="13">
          <el-button @click="onSubmit" type="primary">立即创建</el-button>
          <el-button>取消</el-button>
        </el-col>
        
      </el-row>

      <!-- <el-form-item label="活动名称">
        <el-input v-model="form.name"></el-input>
      </el-form-item>
      <el-form-item>
        <el-select label="活动区域" placeholder="请选择活动区域" v-model="form.region">
          <el-option label="上海" value="shanghai"></el-option>
          <el-option label="北京" value="beijing"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="活动时间">
        <el-col :span="11">
          <el-date-picker placeholder="选择日期" style="width: 100%;" type="date" v-model="form.date1"></el-date-picker>
        </el-col>
        <el-col :span="2" class="line" style="text-align: center">——</el-col>
        <el-col :span="11">
          <el-time-picker placeholder="选择时间" style="width: 100%;" v-model="form.date2"></el-time-picker>
        </el-col>
      </el-form-item>
      <el-form-item label="即时配送">
        <el-switch v-model="form.delivery"></el-switch>
      </el-form-item>
      <el-form-item label="活动性质">
        <el-checkbox-group v-model="form.type">
          <el-checkbox label="美食/餐厅线上活动" name="type"></el-checkbox>
          <el-checkbox label="地推活动" name="type"></el-checkbox>
          <el-checkbox label="线下主题活动" name="type"></el-checkbox>
          <el-checkbox label="单纯品牌曝光" name="type"></el-checkbox>
        </el-checkbox-group>
      </el-form-item>
      <el-form-item label="特殊资源">
        <el-radio-group v-model="form.resource">
          <el-radio label="线上品牌商赞助"></el-radio>
          <el-radio label="线下场地免费"></el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item label="活动形式">
        <el-input type="textarea" v-model="form.desc"></el-input>
      </el-form-item>
      <el-form-item label="穿梭框">
        <el-transfer
          :data="data"
          :filter-method="filterMethod"
          filter-placeholder="请输入城市拼音"
          filterable
          v-model="value"
        ></el-transfer>
      </el-form-item>
      <el-form-item>
        <el-button @click="onSubmit" type="primary">立即创建</el-button>
        <el-button>取消</el-button>
      </el-form-item> -->
    </el-form>
  </div>
</template>
<script>
export default {
  data() {
    const generateData = () => {
      const data = []
      const cities = ['上海', '北京', '广州', '深圳', '南京', '西安', '成都']
      const pinyin = [
        'shanghai',
        'beijing',
        'guangzhou',
        'shenzhen',
        'nanjing',
        'xian',
        'chengdu'
      ]
      cities.forEach((city, index) => {
        data.push({
          label: city,
          key: index,
          pinyin: pinyin[index]
        })
      })
      return data
    }
    return {
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      },
      data: generateData(),
      value: [],
     
    }
  },
  methods: {
       filterMethod(query, item) {
        return item.pinyin.indexOf(query) > -1
      },
    onSubmit() {
      this.$message({
        message: '创建成功',
        type: 'success'
      })
    }
  }
}
</script>
