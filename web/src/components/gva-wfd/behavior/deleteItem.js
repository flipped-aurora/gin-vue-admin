/* eslint-disable */
export default function(G6) {
    G6.registerBehavior('deleteItem', {
        getEvents() {
            return {
                'keydown': 'onKeydown',
                'canvas:mouseleave': 'onCanvasLeave',
                'canvas:mouseenter': 'onCanvasFocus',
            }
        },
        onKeydown(e) {
            const items = this.graph.get('selectedItems');
            const focus = this.graph.get('focusGraphWrapper');
            console.log(e.keyCode)
            if (e.keyCode === 46 && items && items.length > 0 && focus) {
                if (this.graph.executeCommand) {
                    this.graph.executeCommand('delete', {});
                } else {
                    this.graph.remove(items[0]);
                }
                this.graph.set('selectedItems', []);
                this.graph.emit('afteritemselected', []);
            }
            if (e.ctrlKey == true && e.keyCode == 90) { //Ctrl+z
                e.returnvalue = false;
                if (this.graph.executeCommand) {
                    this.graph.executeCommand('undo', {});
                }
            }
            if (e.ctrlKey == true && e.keyCode == 89) { //Ctrl+y
                e.returnvalue = false;
                if (this.graph.executeCommand) {
                    this.graph.executeCommand('redo', {});
                }
            }
            if (e.ctrlKey == true && e.keyCode == 67) { //Ctrl+c
                e.returnvalue = false;
                if (this.graph.executeCommand) {
                    this.graph.executeCommand('copy', {});
                }
            }
            if (e.ctrlKey == true && e.keyCode == 86) { //Ctrl+v
                e.returnvalue = false;
                if (this.graph.executeCommand) {
                    this.graph.executeCommand('paste', {});
                }
            }
        },
        onCanvasLeave(e) {
            this.graph.set('focusGraphWrapper', false);
        },
        onCanvasFocus() {
            this.graph.set('focusGraphWrapper', true);
        }
    });
}