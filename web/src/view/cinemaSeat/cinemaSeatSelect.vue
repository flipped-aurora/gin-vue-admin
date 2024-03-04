<template>
  <div className="container">
  <div class="choose-seat-div">
    <div class="screen">银幕(Screen)</div>
    <div :key="indexRow" class="row-div" v-for="(row,indexRow) in seatMap">
      <div class="row-number">{{indexRow + 1}}排</div>
      <el-row>
        <div  v-for="(col, indexCol) in row" :key="`${indexRow}${indexCol}`">
          <div v-if="col.type !== 'couple'" @click="() => handleClickSingle(indexRow, indexCol)">
            <el-col>
              <SingleSeat
                :name="String(col.number)"
                :status="String(col.status)"
              ></SingleSeat>
            </el-col>
          </div>
          <div v-else  @click="() => handleClickSingle(indexRow, indexCol)">
            <el-col>
              <CoupleSeat
                :name="col.number"
                :status="col.status"
              ></CoupleSeat>
            </el-col>
          </div>
        </div>
      </el-row>
    </div>
  </div>
  <div class="validate-info">
        <div :style="{ marginBottom: '10px'}">您已选择了:</div>
        <div :style="{ display: 'flex' }">
          <div class='print-box' v-for="item in selectedSeat">
                <el-tag color="volcano">
                  <span class="seat-info">
                    {{item[0]}} 排 {{item[1]}} 座
                  </span>
                </el-tag>
              </div>
        </div>
      </div>
      <EasyPrint :printList="selectedSeat" :hall="hall" :info="info" @handleAfterPrint="handleAfterPrint" />
    </div>
</template>
<script lang='ts' setup>
import { ref, watch } from 'vue'
import { ElMessageBox as Modal, ElMessage as message } from 'element-plus'
import { storage } from '@/utils/storage';
import {
  hallSeat1,
  hallSeat2,
  hallSeat3,
  hallSeat4,
  hallSeat5,
  hallSeat6,
  hallSeat7,
} from './contant';

import SingleSeat from './../../components/singleSeat/index.vue'
import CoupleSeat from './../../components/coupleSeat/index.vue'
import EasyPrint from './../../components/easyPrint/index.vue'
const { confirm } = Modal

const props=defineProps<{
  filmId:any
  hallId:any
  seatInfo: any
  propFilmOptions: any
}>()

const hall = ref(1)
const film = ref('')
const seatMap = ref([])
const selectedSeat = ref([])
const filmOptions = ref([])
const totalNum = ref([])
const c = ref(1)
const info = ref({})
  const seat = {
    1: hallSeat1,
    2: hallSeat2,
    3: hallSeat3,
    4: hallSeat4,
    5: hallSeat5,
    6: hallSeat6,
    7: hallSeat7,
  };
const setHall = (val)=> hall.value = val
const setFilm = (val)=> film.value = val
const setSeatMap = (val)=> seatMap.value = val
const setSelectedSeats = (val)=> selectedSeat.value = val
const setFilmOptions = (val)=> filmOptions.value = val
const setTotalNum = (val)=> totalNum.value = val
const setC = (val)=> c.value = val
const setInfo = (val)=> info.value = val

let tempTotalNum = parseInt(storage.get("totalPriceNum"));

const emits = defineEmits(['printSeatSave', 'printSeatDel'])

watch(()=>props.hallId,(v)=>{
  setHall(v)
}, {immediate:true})
watch(()=>props.filmId,(v)=>{
  setFilm(v)
}, {immediate:true})
watch(()=>props.propFilmOptions,(v)=>{
  setFilmOptions(v)
}, {immediate:true})
watch(()=>props.seatInfo,(v)=>{
   console.log(props.seatInfo)
   const tempMap = seat[props.hallId]
   for (let i = 0; i < props.seatInfo.length; i++) {
    const [row, col] = props.seatInfo[i].split('-')
    tempMap[row - 1][col - 1].status = 3
   }
   setSeatMap(tempMap)
}, {immediate:true})

const handleClickSingle = (indexRow: number, indexCol: number) => {
    if (!hall.value || !film.value) {
      message.info('请先选择电影场次！');
      return;
    }

    const tempMap = JSON.parse(JSON.stringify(seatMap.value));
    if (tempMap[indexRow][indexCol].status === 1) {
      tempMap[indexRow][indexCol].status = 2;
      setSeatMap(tempMap);
      return;
    }
    if (tempMap[indexRow][indexCol].status === 2) {
      tempMap[indexRow][indexCol].status = 1;
      setSeatMap(tempMap);
      return;
    }

    if (tempMap[indexRow][indexCol].status === 3) {
      confirm(
        '该座位已出票，确认取消锁定该座位吗？',
        '提示',
      ).then(()=>{
        tempMap[indexRow][indexCol].status = 1;
          setSeatMap(tempMap);
          emits('printSeatDel', tempMap[indexRow][indexCol])
      });
    }
  };

  
  const handleAfterPrint = (seats: any[]) => {
    var postions = seats.map((item) => {
      return item[0] +  "-"  + item[1];
    });
    emits('printSeatSave', postions)
    let tempMap = JSON.parse(JSON.stringify(seatMap.value));
    seats.forEach((seat) => {
      tempMap = tempMap.map((row, rowIndex) =>
        row.map((col, colIndex) => {
          const isCoupleSeat = col.status === 2 && col.type === 'couple';
          const isSingleSeat = col.status === 2 && col.type === 'single';
          const isMatchingCoupleSeat = isCoupleSeat && rowIndex === seat[0] - 1 && col.number.includes(seat[1]);
          const isMatchingSingleSeat = isSingleSeat && rowIndex === seat[0] - 1 && col.number === seat[1];
          if (isMatchingCoupleSeat || isMatchingSingleSeat) {
            return { ...col, status: 3 };
          } else {
            return col;
          }
        })
      );
    });
    setSeatMap(tempMap);
  }

  // watch(()=>c.value,()=>{
  //   const storageList = storage.get('movieList');
  //   if (!storageList) return;
  //   const movieList = JSON.parse(storageList);
  //   const fOptions: Record<string, any[]> = {};
  //   movieList.forEach((item: any) => {
  //     if (fOptions[item.hall]) {
  //       fOptions[item.hall].push({
  //         value: item.key,
  //         label: `${item.name}  (${item.time})`,
  //         info: item,
  //       });
  //     } else {
  //       fOptions[item.hall] = [
  //         {
  //           value: item.key,
  //           label: `${item.name}  (${item.time})`,
  //           info: item,
  //         },
  //       ];
  //     }
  //   });
  //   setFilmOptions(fOptions);
  // }, {immediate:true})

  watch(()=>seatMap.value,()=>{
    const tempSelected: any[] = [];
    seatMap.value.forEach((row, rowIndex) => {
      row.forEach((col) => {
        if (col.status === 2 && col.type === 'single') {
          tempSelected.push([rowIndex + 1, col.number]);
        }
        if (col.status === 2 && col.type === 'couple') {
          const [s1, s2] = col.number;
          tempSelected.push([rowIndex + 1, s1], [rowIndex + 1, s2]);
        }
      });
    });
    setSelectedSeats(tempSelected);
  })

  watch(()=>[film.value, c.value],()=>{
    if (!hall.value) return;
    if (film.value === '') {
      setSeatMap(seat[hall.value]);
    } else {
      if (filmOptions.value?.length) {
      setInfo(
        filmOptions.value?.filter((item) => item.value === film.value)?.[0].info,
      );
      const hallSeat = JSON.parse(storage.get(film.value));
      setSeatMap(hallSeat);
      }
    }
  },{immediate:true})
</script>
<style lang='scss' scoped>
.container {
  width: 100%;
  height: 100%;
  padding: 16px;

  .select-div {
    display: flex;
    justify-content: space-between;
    margin-bottom: 30px;
  }

  .top-bar {
    display: flex;
    justify-content: flex-end;
  }

  .choose-seat-div {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding: 4px 10px 16px 10px;

    .row-div {
      position: relative;

      .row-number {
        height: 30px;
        line-height: 30px;
        position: absolute;
        left: -50px;
      }
    }

    .screen {
      background-color: silver;
      width: 200px;
      height: 20px;
      font-size: 16px;
      font-family: Verdana, Geneva, Tahoma, sans-serif;
      border-radius: 10px;
      text-align: center;
      margin-bottom: 12px;
    }
  }

  .validate-info {
    .print-box {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      width: 94px;
      margin-right: 10px;

      .seat-info {
        padding: 0px 12px;
      }
    }

  }
}
</style>