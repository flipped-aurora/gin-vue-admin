export default function(G6){
  G6.registerBehavior('deleteItem', {
    getEvents() {
      return {
        'keydown': 'onKeydown',
        'canvas:mouseleave': 'onCanvasLeave',
        'canvas:mouseenter': 'onCanvasFocus',
      }
    },
    onKeydown(e){
      const items = this.graph.get('selectedItems');
      const focus = this.graph.get('focusGraphWrapper');
      if(e.keyCode === 8 && items && items.length > 0 && focus){
        if(this.graph.executeCommand) {
          this.graph.executeCommand('delete', {});
        }else{
          this.graph.remove(items[0]);
        }
        this.graph.set('selectedItems',[]);
        this.graph.emit('afteritemselected',[]);
      }
    },
    onCanvasLeave(e){
      this.graph.set('focusGraphWrapper',false);
    },
    onCanvasFocus(){
      this.graph.set('focusGraphWrapper',true);
    }
  });
}
