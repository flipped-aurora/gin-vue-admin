export default function(G6){
  G6.registerBehavior('clickSelected', {
    getDefaultCfg() {
      return {
        multiple: false,
      }
    },
    getEvents() {
      return {
        'node:click': 'onClick',
        'edge:click': 'onClick',
        'edge:mouseover': 'onEdgeMouseOver',
        'edge:mouseleave': 'onEdgeMouseLeave',
        'canvas:click': 'onCanvasClick',
        'node:mouseover': 'onNodeMouseOver',
      }
    },
    onClick(e) {
      this._clearSelected();
      this.graph.setItemState(e.item, 'selected', true);
      let selectedItems = this.graph.get('selectedItems');
      if(!selectedItems)
        selectedItems = [];
      selectedItems = [e.item.get('id')];
      this.graph.set('selectedItems',selectedItems);
      this.graph.emit('afteritemselected',selectedItems);
    },
    onNodeMouseOver(e){
      if(this.graph.getCurrentMode() === 'edit')
        this.graph.setItemState(e.item, 'hover', true);
      else
        this.graph.setItemState(e.item, 'hover', false);
    },
    onEdgeMouseOver(e){
      if(this.graph.getCurrentMode() === 'edit' && !e.item.hasState('selected'))
        this.graph.setItemState(e.item, 'hover', true);
    },
    onEdgeMouseLeave(e){
      if(this.graph.getCurrentMode() === 'edit' && !e.item.hasState('selected'))
        this.graph.setItemState(e.item, 'hover', false);
    },
    onCanvasClick(){
      this._clearSelected();
      this.graph.emit('afteritemselected',[]);
    },
    _clearSubProcessSelected() {
      const subProcessList = this.graph.findAll('node', (node) => {
        if (node.get('model')) {
          const clazz = node.get('model').clazz;
          return clazz === 'subProcess';
        } else {
          return false;
        }
      });
      subProcessList.forEach((node) => {
        const group = node.getContainer();
        const subGroup = group.subGroup;
        this._clearGroupSelected(subGroup);
      });
    },
    _clearGroupSelected(group) {
      const selected = group.findAll((subGroup) => {
        const node = subGroup.get('item');
        if (node) {
          return node.hasState('selected');
        } else {
          return false;
        }
      });
      selected.forEach(subGroup => {
        const node = subGroup.get('item');
        if (node) {
          node.setState('selected', false);
        }
      });
    },
    _clearSelected(){
      let selected = this.graph.findAllByState('node', 'selected');
      selected.forEach(node => {
        this.graph.setItemState(node, 'selected', false);
      });
      selected = this.graph.findAllByState('edge', 'selected');
      selected.forEach(edge => {
        this.graph.setItemState(edge, 'selected', false);
      });
      this._clearSubProcessSelected();
      this.graph.set('selectedItems', []);
    }
  });
}
