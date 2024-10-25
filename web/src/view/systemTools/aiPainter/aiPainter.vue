<template>
<div>
    <div class="w-full">
    </div>
    <div class="gva-search-box" >
        <div class="w-full relative">
        <el-input v-model="prompt" type="textarea" :rows="8" class="w-full"></el-input>
        <el-button :loading="loading" @click="aiPainter" type="primary" class="absolute right-2 bottom-2"><ai-gva />{{html?'修改':'生成'}}</el-button>
        </div>
    </div>
</div>
</template>

<script setup>
import {painter} from "@/api/autoCode";
import { ref } from 'vue'
const loading = ref(false)
const prompt = ref('')
const html = ref(``)
const aiPainter = async () => {
    const res = await painter({
        html: html.value,
        prompt: prompt.value
    })
    html.value = res.data
}
</script>