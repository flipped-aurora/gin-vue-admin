export default function(G6){
  G6.registerBehavior('hoverAnchorActived', {
    getEvents() {
      return {
        'anchor:mouseenter': 'onAnchorEnter',
        'anchor:mousemove': 'onAnchorEnter',
        'anchor:mouseleave': 'onAnchorLeave',
      }
    },
    onAnchorEnter(e){
      if(!this.graph.get('edgeDragging'))
        this.graph.setItemState(e.item, 'active-anchor', true);
    },
    onAnchorLeave(e){
      if(!this.graph.get('edgeDragging')) {
        let node = e.item.getContainer().getParent();
        if(node) {
          this.graph.setItemState(e.item, 'active-anchor', false);
        }
      }
    }
  });
}
