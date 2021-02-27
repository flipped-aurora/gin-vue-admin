<template>
    <div class="previewCode">
        <el-tabs v-model="activeName">
            <el-tab-pane :label="key" :name="key" v-for="(item,key) in previewCode" :key="key">
                <div style="background:#fff;padding:0 20px" :id="key"></div>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>

<script>
import marked from "marked"
import hljs from "highlight.js";
// import 'highlight.js/styles/atelier-cave-light.css';
import 'highlight.js/styles/atelier-plateau-light.css';
export default {
    props:{
        previewCode:{
            type:Object,
            default(){
                return {}
            }
        }
    },
    data(){
        return{
             activeName: "",
        }
    },
    mounted(){
        marked.setOptions({
          renderer: new marked.Renderer(),
          highlight: function(code) {
            return hljs.highlightAuto(code).value;
          },
          pedantic: false,
          gfm: true,
          tables: true,
          breaks: false,
          sanitize: false,
          smartLists: true,
          smartypants: false,
          xhtml: false
        }
      );
        for(const key in this.previewCode){
            if(this.activeName == ""){
                this.activeName = key
            }
            document.getElementById(key).innerHTML = marked(this.previewCode[key])
        }
    }
}
</script>

<style lang="scss">
    
</style>