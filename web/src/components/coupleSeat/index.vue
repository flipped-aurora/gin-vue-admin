<template>
  <el-row>
    <el-col>
        <div :class="`couple-seat-${status}`">
          <span>{{ name[0] }}</span>
          <span>{{ name[1] }}</span>  
        </div>
      </el-col>
  </el-row>

</template>
<script lang='ts' setup>
interface Props {
  name: Array<number>;
  status: string;
}
const props=defineProps<Props>()

</script>
<style lang='scss' scoped>
.couple-seat-1 {
  cursor: pointer;
  width: 64px;
  height: 30px;
  padding: 2px;
  margin: 2px;
  border: 0.5px solid rgba(0, 0, 0, 0.18);
  border-radius: 6px;
  display: flex;
  justify-content: center;
  align-items: center;

  span {
    background-image: url('../../assets/empty.png');
    display: flex;
    justify-content: center;
    align-items: center;
    width: 30px;
    height: 26px;
  }
}

.couple-seat-3 {
  width: 64px;
  height: 30px;
  padding: 2px;
  margin: 2px;
  border: 0.5px solid rgba(0, 0, 0, 0.18);
  border-radius: 6px;
  display: flex;
  justify-content: center;
  align-items: center;

  span {
    background-image: url('../../assets/reserved.png');
    display: flex;
    justify-content: center;
    align-items: center;
    width: 30px;
    height: 26px;
  }
}

.single-seat-0 {
  width: 30px;
  height: 26px;
  margin: 2px;
  visibility: hidden;
}

.couple-seat-2 {
  cursor: pointer;
  width: 64px;
  height: 30px;
  padding: 2px;
  margin: 2px;
  border: 0.5px solid rgba(0, 0, 0, 0.18);
  border-radius: 6px;
  display: flex;
  justify-content: center;
  align-items: center;

  span {
    display: flex;
    justify-content: center;
    align-items: center;
    color: transparent;
    width: 30px;
    height: 26px;
    background-image: url('../../assets/choose.png');
  }
}
</style>