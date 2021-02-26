<template>
  <div class="root">
    <ToolbarPanel ref="toolbar" v-if="!isView" />
    <div style="display: flex">
      <ItemPanel ref="addItemPanel" v-if="!isView" :height="height"/>
      <div ref="canvas" class="canvasPanel" :style="{'height':height+'px','width':isView?'100%':'70%','border-bottom':isView?0:null}"></div>
      <DetailPanel ref="detailPanel"
                   v-if="!isView"
                   :height="height"
                   :model="selectedModel"
                   :readOnly="$route.query.type == 'view'"
                   :users="users"
                   :authorities="authorities"
                   :groups="groups"
                   :categorys="categorys"
                   :signalDefs="processModel.signalDefs"
                   :messageDefs="processModel.messageDefs"
                   :onChange="(key,val)=>{onItemCfgChange(key,val)}" />
    </div>
  </div>
</template>
<script>
  import G6 from '@antv/g6/lib';
  import { getShapeName } from '../util/clazz'
  import Command from '../plugins/command'
  import Toolbar from '../plugins/toolbar'
  import AddItemPanel from '../plugins/addItemPanel'
  import CanvasPanel from '../plugins/canvasPanel'
  import ToolbarPanel from '../components/ToolbarPanel'
  import ItemPanel from '../components/ItemPanel'
  import DetailPanel from '../components/DetailPanel'
  import i18n from '../locales'
  import {exportXML,exportImg} from "../util/bpmn"
  import registerShape from '../shape'
  import registerBehavior from '../behavior'
  registerShape(G6);
  registerBehavior(G6);
  export default {
    name: "wfd-vue",
    components: {
      ToolbarPanel,
      ItemPanel,
      DetailPanel
    },
    provide() {
      return {
        i18n: i18n[this.lang]
      }
    },
    props: {
      isView: {
        type: Boolean,
        default: false,
      },
      mode: {
        type: String,
        default: "edit"
      },
      height: {
        type: Number,
        default: 800,
      },
      lang: {
        type: String,
        default: "zh"
      },
      data: {
        type: Object,
        default: () => ({nodes:[],edges:[]})
      },
      users: {
        type: Array,
        default: () => ([])
      },
      authorities: {
        type: Array,
        default: ()=>([]),
      },
      groups: {
        type: Array,
        default: () => ([])
      },
      categorys: {
        type: Array,
        default: () => ([])
      },
      propProcessModel:{
         type: Object,
          default: () => ({
          id: '',
          name: '',
          category: '',
          clazz: 'process',
          dataObjs: [],
          signalDefs: [],
          messageDefs: [],
        })
      }
    },
    data() {
      return {
        processModel:{},
        resizeFunc: ()=>{},
        selectedModel: {},
        graph:null,
        cmdPlugin: null,
      };
    },
    watch:{
      data(oldData,newData){
        if(oldData !== newData) {
          if (this.graph) {
            this.graph.changeData(this.initShape(newData));
            this.graph.setMode(this.mode);
            this.graph.emit('canvas:click');
            if (this.cmdPlugin) {
              this.cmdPlugin.initPlugin(this.graph);
            }
            if (this.isView) {
              this.graph.fitView(5)
            }
          }
        }
      },
    },
    methods: {
      initShape(data){
        if(data && data.nodes){
          return {
            nodes: data.nodes.map(node => {
              return {
                shape: getShapeName(node.clazz),
                ...node,
              }
            }),
            edges: data.edges
          }
        }
        return data;
      },
      initEvents(){
        this.graph.on('afteritemselected',(items)=>{
          if(items && items.length > 0) {
            let item = this.graph.findById(items[0]);
            if(!item){
              item = this.getNodeInSubProcess(items[0])
            }
            this.selectedModel = {...item.getModel()};
          } else {
            this.selectedModel = this.processModel;
          }
        });
         this.graph.on('aftercommandexecute',(command)=>{
          if (command.name !== 'undo' && command.name !== 'redo') return;
            // 更新selectedModel
         });
        const page = this.$refs['canvas'];
        const graph = this.graph;
        const height = this.height-1;
        this.resizeFunc = ()=>{
          graph.changeSize(page.offsetWidth,height);
        };
        window.addEventListener("resize", this.resizeFunc);
      },
      onItemCfgChange(key,value){
        const items = this.graph.get('selectedItems');
        if(items && items.length > 0){
          let item = this.graph.findById(items[0]);
          if(!item){
            item = this.getNodeInSubProcess(items[0])
          }
          if(this.graph.executeCommand) {
            this.graph.executeCommand('update', {
              itemId: items[0],
              updateModel: {[key]: value}
            });
          }else {
            this.graph.updateItem(item, {[key]: value});
          }
          this.selectedModel = {...item.getModel()};
        } else {
          const canvasModel = { ...this.processModel, [key]: value};
          this.selectedModel = canvasModel;
          this.processModel = canvasModel;
        }
      },
      getNodeInSubProcess(itemId){
        const subProcess = this.graph.find('node', (node) => {
          if (node.get('model')) {
            const clazz = node.get('model').clazz;
            if (clazz === 'subProcess') {
              const containerGroup = node.getContainer();
              const subGroup = containerGroup.subGroup;
              const item = subGroup.findById(itemId);
              return subGroup.contain(item);
            } else {
              return false;
            }
          } else {
            return false;
          }
        });
        if(subProcess) {
          const group = subProcess.getContainer();
          return group.getItem(subProcess, itemId);
        }
        return null;
      },
    },
    destroyed(){
      window.removeEventListener("resize", this.resizeFunc);
      this.graph.getNodes().forEach(node => {
        node.getKeyShape().stopAnimate();
      });
    },
    mounted() {
      let plugins = [];
      if(!this.isView){
        this.cmdPlugin = new Command();
        const toolbar = new Toolbar({container:this.$refs['toolbar'].$el});
        const addItemPanel = new AddItemPanel({container:this.$refs['addItemPanel'].$el});
        const canvasPanel = new CanvasPanel({container:this.$refs['canvas']});
        plugins = [ this.cmdPlugin,toolbar,addItemPanel,canvasPanel ];
      }
      const width = this.$refs['canvas'].offsetWidth;
      this.graph = new G6.Graph({
        plugins: plugins,
        container: this.$refs['canvas'],
        height: this.height,
        width: width,
        modes: {
          default: ['drag-canvas', 'clickSelected'],
          view: [ ],
          edit: [ 'drag-canvas', 'hoverNodeActived','hoverAnchorActived','dragNode','dragEdge',
            'dragPanelItemAddNode','clickSelected','deleteItem','itemAlign','dragPoint','brush-select'],
        },
        defaultEdge: {
          shape: 'flow-polyline-round',
        },
      });
      this.graph.saveXML = (createFile = true) => exportXML(this.graph.save(),this.processModel,createFile);
      this.graph.saveImg = (createFile = true) => exportImg(this.$refs['canvas'],this.processModel.name,createFile);
      if(this.isView)
        this.graph.setMode('view');
      else
        this.graph.setMode(this.mode);
      this.graph.data(this.initShape(this.data));
      this.graph.render();
      if(this.isView && this.data && this.data.nodes){
        this.graph.fitView(5)
      }
      this.initEvents();
    },
    created() {
      this.processModel = this.propProcessModel
      this.selectedModel = this.processModel
    },
  };
</script>
<style lang="scss" scoped>
    .root{
        width: 100%;
        height: 100%;
        background-color: #fff;
        display: block;
    }
    .canvasPanel {
        flex: 0 0 auto;
        float: left;
        width:70%;
        background-color: #fff;
        border-bottom: 1px solid #E9E9E9;
    }
</style>
