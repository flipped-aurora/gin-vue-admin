<template>
  <div class="container">
    <ul class="moon">
      <li class="crater"></li>
      <li class="crater"></li>
      <li class="crater"></li>
    </ul>
    <ul class="mountain-range">
      <li class="mountain"></li>
      <li class="mountain"></li>
      <li class="mountain"></li>
      <li class="mountain"></li>
      <li class="mountain"></li>
      <li class="mountain"></li>
      <li class="mountain"></li>
      <li class="mountain"></li>
      <li class="mountain"></li>
      <li class="mountain"></li>
      <li class="mountain"></li>
    </ul>
    <div class="mountain-range-mask"></div>
    <ul class="forest">
      <li class="hill"></li>
      <li class="hill"></li>
      <li class="hill"></li>
    </ul>
    <ul class="sparkles">
      <li class="sparkle one"></li>
      <li class="sparkle one"></li>
      <li class="sparkle one"></li>
      <li class="sparkle one"></li>
      <li class="sparkle two"></li>
      <li class="sparkle two"></li>
      <li class="sparkle two"></li>
      <li class="sparkle two"></li>
      <li class="sparkle three"></li>
      <li class="sparkle three"></li>
      <li class="sparkle three"></li>
      <li class="sparkle three"></li>
      <li class="sparkle four"></li>
      <li class="sparkle four"></li>
      <li class="sparkle four"></li>
      <li class="sparkle four"></li>
      <li class="sparkle five"></li>
      <li class="sparkle five"></li>
      <li class="sparkle five"></li>
      <li class="sparkle five"></li>
      <li class="sparkle six"></li>
      <li class="sparkle six"></li>
      <li class="sparkle six"></li>
      <li class="sparkle six"></li>
      <li class="sparkle seven"></li>
      <li class="sparkle seven"></li>
      <li class="sparkle seven"></li>
      <li class="sparkle seven"></li>
      <li class="sparkle eight"></li>
      <li class="sparkle eight"></li>
      <li class="sparkle eight"></li>
      <li class="sparkle eight"></li>
    </ul>
    <div class="grass">
      <div class="pokemon">
        <div class="poke" id="bulbasaur">
          <div class="ear"></div>
          <div class="ear"></div>
          <div class="head"></div>
          <div class="leg"></div>
          <div class="bulba-body"></div>
          <div class="bulbs">
            <div class="bulb"></div>
          </div>
        </div>
        <div class="poke" id="pikachu">
          <div class="ear"></div>
          <div class="ear"></div>
          <div class="hand"></div>
          <div class="pika-body"></div>
          <div class="head"></div>
          <div class="pika-tail"></div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'Animition'
}
</script>
<style lang="scss">
  @import '@/style/animition.scss';

</style>