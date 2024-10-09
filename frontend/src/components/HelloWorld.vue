<script setup>
import {onMounted, reactive} from 'vue'
import {Greet, SetWallpaper} from '../../wailsjs/go/main/App'

const data = reactive({
  wallpaper: "",
  wallpapers: [],
  currenIndex: -1,
})


function prevWallpaper() {
  if (data.currenIndex === 0) {
    return
  }
  data.currenIndex -= 1
  data.wallpaper = data.wallpapers[data.currenIndex]
}

function nextWallpaper() {
  if (data.currenIndex < data.wallpapers.length - 1) {
    data.currenIndex += 1
    data.wallpaper = data.wallpapers[data.currenIndex]
  } else  {
    Greet().then(result => {
      data.wallpaper = result;
      data.currenIndex ++;
      data.wallpapers.push(result)
    })
  }
}

function setWallpaper() {
  SetWallpaper(data.wallpaper);
}

onMounted(() => {
  nextWallpaper();
});

</script>

<template>
  <img id="logo" alt="Wails logo" :src="data.wallpaper" @load="data.wallpaper"/>
  <main>
    <div id="input" class="input-box">
      <button class="btn" @click="prevWallpaper" :disabled="data.currenIndex === 0">上一张</button>
      <button class="btn" @click="nextWallpaper">下一张</button>
      <button class="btn" @click="setWallpaper">设为背景</button>
    </div>
  </main>
</template>

<style scoped>
#logo {
  display: block;
  width: 95%;
  height: 88%;
  margin: auto;
  padding: 1.5% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}


.input-box .btn {
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 20px 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}
</style>
