<template>
  <div>
    <div class="big">
      <div class="inner">
        <img src="../../assets/notFound.png">
        <p>页面被神秘力量吸走了，请联系我们修复</p>
        <p style="font-size:18px;line-height:40px;">常见问题为菜单未分配仪表盘，可自行分配仪表盘或修改默认路由</p>
        <p>↓</p>
        <img src="../../assets/qm.png" class="leftPic">
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Error'

}
</script>

<style lang="scss">
    .big{
        width: 100%;
        height: calc(100vh - 220px);
        background-color: rgb(244,244,244);
        position: relative;
    }
    .inner{
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
    .inner p {
        text-align: center;
        font-size: 24px;
    }
    .inner .leftPic{
        width: 60px;
        height: 60px;
        margin-left: 44%;
        margin-top: 20px;
    }
</style>
