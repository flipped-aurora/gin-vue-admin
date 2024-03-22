<template>
    <div class="tinymce-box">
		<Editor v-model="contentValue" :init="init" :disabled="disabled" @onClick="onClick" />
	</div>
      <ChooseImg ref="chooseImg" @enterImg="enterImg" />
</template>

<script setup>

import { ref, shallowRef, onMounted, onBeforeUnmount, nextTick,watch } from 'vue'
const path = ref(import.meta.env.VITE_BASE_API)
const serverpath = ref(import.meta.env.VITE_BASE_PATH + ":" + import.meta.env.VITE_SERVER_PORT + "/")

import service from '@/utils/request'
//引入tinymce编辑器
import Editor from '@tinymce/tinymce-vue'
import ChooseImg from '@/components/chooseImg/index.vue'
//引入node_modules里的tinymce相关文件文件
import tinymce from 'tinymce/tinymce' //tinymce默认hidden，不引入则不显示编辑器
import 'tinymce/themes/silver'  //编辑器主题，不引入则报错
import 'tinymce/icons/default'  //引入编辑器图标icon，不引入则不显示对应图标

// 引入编辑器插件（基本免费插件都在这儿了）
import 'tinymce/plugins/advlist'  //高级列表
import 'tinymce/plugins/anchor'  //锚点
import 'tinymce/plugins/autolink'  //自动链接
import 'tinymce/plugins/autoresize'  //编辑器高度自适应,注：plugins里引入此插件时，Init里设置的height将失效
import 'tinymce/plugins/autosave'  //自动存稿
import 'tinymce/plugins/charmap'  //特殊字符
import 'tinymce/plugins/code'  //编辑源码
import 'tinymce/plugins/codesample'  //代码示例
import 'tinymce/plugins/directionality'  //文字方向
import 'tinymce/plugins/emoticons'  //表情
import 'tinymce/plugins/fullpage'  //文档属性
import 'tinymce/plugins/fullscreen'  //全屏
import 'tinymce/plugins/help'  //帮助
import 'tinymce/plugins/hr'  //水平分割线
import 'tinymce/plugins/image'  //插入编辑图片
import 'tinymce/plugins/importcss'  //引入css
import 'tinymce/plugins/insertdatetime'  //插入日期时间
import 'tinymce/plugins/link'  //超链接
import 'tinymce/plugins/lists' //列表插件
import 'tinymce/plugins/media' //插入编辑媒体
import 'tinymce/plugins/nonbreaking' //插入不间断空格
import 'tinymce/plugins/pagebreak' //插入分页符
import 'tinymce/plugins/paste' //粘贴插件
import 'tinymce/plugins/preview'//预览
import 'tinymce/plugins/print'//打印
import 'tinymce/plugins/quickbars'  //快速工具栏
import 'tinymce/plugins/save'  //保存
import 'tinymce/plugins/searchreplace'  //查找替换
// import 'tinymce/plugins/spellchecker'  //拼写检查，暂未加入汉化，不建议使用
import 'tinymce/plugins/tabfocus'  //切入切出，按tab键切出编辑器，切入页面其他输入框中
import 'tinymce/plugins/table'  //表格
import 'tinymce/plugins/template'  //内容模板
import 'tinymce/plugins/textcolor'  //文字颜色
import 'tinymce/plugins/textpattern'  //快速排版
import 'tinymce/plugins/toc'  //目录生成器
import 'tinymce/plugins/visualblocks'  //显示元素范围
import 'tinymce/plugins/visualchars'  //显示不可见字符
import 'tinymce/plugins/wordcount'  //字数统计
// import 'tinymce/plugins/axupimgs'  //多图片批量上传

const chooseImg = ref(null)

const props = defineProps({
      
      content: {
            type: String,
            default: ''
      },
      disabled: {
            type: Boolean,
            default: false
      },
      plugins: {
            type: [String, Array],
            default: 'print preview searchreplace autolink directionality visualblocks visualchars fullscreen image link media template code codesample table charmap hr pagebreak nonbreaking anchor insertdatetime advlist lists wordcount textpattern autosave'
      },
      toolbar: {
            type: [String, Array], // image axupimgs media
            default: 'fullscreen code preview undo redo restoredraft | table uploadimages hr pagebreak insertdatetime | forecolor backcolor bold italic underline strikethrough link anchor |alignleft aligncenter alignright alignjustify outdent indent | \
            styleselect formatselect fontselect fontsizeselect | bullist numlist | blockquote subscript superscript removeformat | \
            selectall searchreplace visualblocks | indent2em lineheight formatpainter print cut copy paste pastetext charmap'
      },

});

const init = {
            language_url: '/tinymce/langs/zh-Hans.js',  //引入语言包文件
            language: 'zh-Hans',  //语言类型

            skin_url: '/tinymce/skins/ui/oxide',  //皮肤：浅色
            // skin_url: '/tinymce/skins/ui/oxide-dark',//皮肤：暗色

            plugins: props.plugins,  //插件配置
            toolbar: props.toolbar,  //工具栏配置，设为false则隐藏
            menubar: false,  //菜单栏配置，设为false则隐藏，不配置则默认显示全部菜单，也可自定义配置--查看 http://tinymce.ax-z.cn/configure/editor-appearance.php --搜索“自定义菜单”

            fontsize_formats: '12px 14px 16px 18px 20px 22px 24px 28px 32px 36px 48px 56px 72px',  //字体大小
            font_formats: '微软雅黑=Microsoft YaHei,Helvetica Neue,PingFang SC,sans-serif;苹果苹方=PingFang SC,Microsoft YaHei,sans-serif;宋体=simsun,serif;仿宋体=FangSong,serif;黑体=SimHei,sans-serif;Arial=arial,helvetica,sans-serif;Arial Black=arial black,avant garde;Book Antiqua=book antiqua,palatino;',  //字体样式
//自带默认字体：'Andale Mono=andale mono,times;Arial=arial,helvetica,sans-serif;Arial Black=arial black,avant garde;Book Antiqua=book antiqua,palatino;Comic Sans MS=comic sans ms,sans-serif;Courier New=courier new,courier;Georgia=georgia,palatino;Helvetica=helvetica;Impact=impact,chicago;Symbol=symbol;Tahoma=tahoma,arial,helvetica,sans-serif;Terminal=terminal,monaco;Times New Roman=times new roman,times;Trebuchet MS=trebuchet ms,geneva;Verdana=verdana,geneva;Webdings=webdings;Wingdings=wingdings,zapf dingbats'
            lineheight_formats: "0.5 0.8 1 1.2 1.5 1.75 2 2.5 3 4 5",  //行高配置，也可配置成"12px 14px 16px 20px"这种形式

            height: 400,  //注：引入autoresize插件时，此属性失效
            placeholder: '在这里输入文字',
            branding: false,  //tiny技术支持信息是否显示
            resize: 'both',  //编辑器宽高是否可变，false-否,true-高可变，'both'-宽高均可，注意引号
            statusbar: false,  //最下方的元素路径和字数统计那一栏是否显示
            elementpath: false,  //元素路径是否显示

            // relative_urls: false,  //false: tinymce将不再自动将文件路径由绝对转为相对
            // convert_urls: false,  //false: tinymce将不再自动处理文件路径

            content_style: "img {max-width:100%;}",  //直接自定义可编辑区域的css样式
            content_css: '/tinymce/skins/content/default/content.css',  //以css文件方式自定义可编辑区域的css样式，css文件需自己创建并引入
      
            // images_upload_url: '/apib/api-upload/uploadimg',  //后端处理程序的url，建议直接自定义上传函数image_upload_handler，这个就可以不用了
            images_upload_base_path: serverpath.value,  //相对基本路径--关于图片上传建议查看--http://tinymce.ax-z.cn/general/upload-images.php
            // paste_data_images: true,  //图片是否可粘贴
            // images_file_types: 'jpeg,jpg,png,gif,bmp,webp', 
            toolbar_mode:"sliding", // 取值：floating / sliding / scrolling / wrap
            images_upload_handler: (blobInfo, success, failure) => {
                  if(blobInfo.blob().size/1024/1024>2){
                        failure("上传失败，图片大小请控制在 2M 以内")
                  }else{
                        let params=new FormData()
                        params.append('file',blobInfo.blob())

                        service({
                        url: '/fileUploadAndDownload/upload',
                        method: 'post',
                        data: params,
                        }).then(res=>{
                              // console.log(res)
                              if(res.code==0){
                                    let fileurl = path.value + res.data.file.url;
                                    success(fileurl)  //上传成功，在成功函数里填入图片路径
                              }else{
                                    failure("上传失败")
                              }
                        }).catch(()=>{
                              failure("上传出错，服务器开小差了呢")
                        })
                  }
            },
            file_picker_callback: (callback, value, meta) => {
                  let filetype =
                  ".pdf, .txt, .zip, .rar, .7z, .doc, .docx, .xls, .xlsx, .ppt, .pptx, .mp3, .mp4,.mkv, .avi,.wmv, .rmvb,.mov,.mpg,.mpeg,.webm, .jpg, .jpeg, .png, .gif"; //限制文件的上传类型
                  let inputElem = document.createElement("input"); //创建文件选择
                  inputElem.setAttribute("type", "file");
                  inputElem.setAttribute("accept", filetype);
                  inputElem.click();
                  inputElem.onchange = () => {
                  let upurl = path.value + '/fileUploadAndDownload/upload';
                  let file = inputElem.files[0]; //获取文件信息
                  let params = new FormData();
                  if (file.type.slice(0, 5) == "video") {
                        //判断文件类型
                        params.append("file", file);
                        params.append("id", props.objId.newId);
                  } else if (file.type.slice(0, 5) == "image") {
                        params.append("file", file);
                        params.append("id", props.objId.newId);
                  } else {
                        upurl = "/news/attachUpload";
                        params.append("file", file);
                        params.append("siteId", props.objId.siteId);
                        params.append("newsId", props.objId.newId);
                        params.append("attachDesc", "");
                  }
                  if (file.type.slice(0, 5) == "image" && file.size / 1024 / 1024 > 2) {
                        alert("上传失败，图片大小请控制在2M以内");
                  } else if (
                        file.type.slice(0, 5) == "video" &&
                        file.size / 1024 / 1024 > 500
                  ) {
                        alert("上传失败，视频大小请控制在 500M 以内");
                  } else if (file.size / 1024 / 1024 > 10) {
                        alert("上传失败，文件大小请控制在 10M 以内");
                  } else {
                        
                  }
                  };
            },
            setup: (editor) => {
                  editor.ui.registry.addButton('uploadimages', {
                        icon: 'image',
                        tooltip: '上传文件',
                        onAction: () => {
                              chooseImg.value.open()
                              // editor.execCommand('CodeSample') // 执行命令
                        }
                  });
            },
      }
const contentValue = ref("")
// 转换图片链接 加上 /api/ 用于后端显示
contentValue.value =  props.content.replaceAll("/uploadsfile/", "api/uploadsfile/")

const emits = defineEmits(['onClick',"update:content"])

watch(contentValue, (newValue, oldValue) => {
      // 转换图片链接 去掉 /api/ 用于前端显示
      emits("update:content",newValue.replaceAll("/api",""))
})
onMounted(() => {
      tinymce.init({})
})

// 添加相关的事件
const onClick = (e) =>{
      emits('onClick', e, tinymce)
}
//清空内容  
const clear = (e) =>{
      contentValue.value = ''
}

const enterImg = (item) => {
  let url = path.value + item.url
  // 插入图片
  if (item.tag == 'png' || item.tag == 'jpg') {
    url = '<img src="' + url + '" alt="' + item.name + '" >'
  }
  else if (item.tag == 'docx' || item.tag == 'doc') {
    url = '<a href="' + url + '" >' + item.name + '</a>'
  }
  else if (item.tag == 'mp4') {
      url = '<video poster="" controls="controls" ><source src="'+url+'" type="video/mp4" /></video>'
  }
  tinymce.execCommand('mceInsertContent', false, url);
}

</script>

<style>
  .tox-tinymce-aux{
    z-index: 130000 !important;
  }
</style>

