<template>
  <el-row>
    <el-col>
        <div :class="`single-seat-${status}`">{{name}}</div>
      </el-col>
  </el-row>

</template>
<script lang='ts' setup>
interface Props {
  name: string;
  status: string;
}
const props=defineProps<Props>()

</script>
<style lang='scss' scoped>
.single-seat-1 {
  cursor: pointer;
  width: 30px;
  height: 26px;
  margin: 2px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-image: url('../../assets/empty.png');
}

.single-seat-3 {
  width: 30px;
  height: 26px;
  margin: 2px;
  display: flex;
  justify-content: center;
  align-items: center;
  background-image: url('../../assets/reserved.png');
}

.single-seat-0 {
  width: 30px;
  height: 26px;
  margin: 2px;
  visibility: hidden;
}

.single-seat-2 {
  cursor: pointer;
  width: 30px;
  height: 26px;
  margin: 2px;
  color: transparent;
  display: flex;
  justify-content: center;
  align-items: center;
  background-image: url('../../assets/choose.png');
}
</style>