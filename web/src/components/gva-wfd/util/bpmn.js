/* eslint-disable */
function tab(len){
  return [...Array(len)].map(a => " ").join('')
}
export function exportXML(json,canvas,createFile = true) {
  const id = canvas.id || "flow";
  const name = canvas.name || "flow";
  let dataObjs = "";
  canvas.dataObjs.forEach(s => {
    dataObjs += `${tab(4)}<dataObject id="${s.id}" name="${s.name}" itemSubjectRef="xsd:${s.type}"></dataObject>\n`;
  });
  let signals = "";
  canvas.signalDefs.forEach(s => {
    signals += `${tab(2)}<signal id="${s.id}" name="${s.name}" flowable:scope="${s.scope}"></signal>\n`;
  });
  let messages = "";
  canvas.messageDefs.forEach(s => {
    messages += `${tab(2)}<message id="${s.id}" name="${s.name}"></message>\n`;
  });
  let BPMNShape = ``;
  let BPMNEdge = ``;

  let processXML = `${tab(2)}<process id="${id}" name="${name}">\n`;
  processXML += dataObjs;
  json.nodes.forEach(node => {
    BPMNShape += `${tab(6)}<bpmndi:BPMNShape bpmnElement="${node.id}" id="BPMNShape_${node.id}">\n`+
      `${tab(8)}<omgdc:Bounds width="${node.size[0]}" height="${node.size[1]}" x="${node.x}" y="${node.y}"></omgdc:Bounds>\n`+
      `${tab(6)}</bpmndi:BPMNShape>\n`;
    switch (node.clazz) {
      case 'start':
        processXML += `${tab(4)}<startEvent id="${node.id}"></startEvent>\n`;
        break;
      case 'end':
        processXML += `${tab(4)}<endEvent id="${node.id}"></endEvent>\n`;
        break;
      case 'userTask': {
        let assignments = "";
        if(node.assignValue && node.assignValue.length > 0){
          if(node.assignType === 'person'){
            assignments += `flowable:candidateUsers="${node.assignValue.join(',')}"`;
          }else if(node.assignType === 'assignee'){
            assignments += `flowable:assignee="${node.assignValue[0]}"`;
          }else if(node.assignType === 'persongroup'){
            assignments += `flowable:candidateGroups="${node.assignValue.join(',')}"`;
          }
        }
        processXML += `${tab(4)}<userTask id="${node.id}" name="${node.label}" ${assignments}></userTask>\n`;
        break;
      }
      case 'javaTask': {
        let javaClass = "";
        if(node.javaClass){
          javaClass = `flowable:class="${node.javaClass}"`;
        }
        processXML += `${tab(4)}<serviceTask id="${node.id}" name="${node.label}" ${javaClass}></serviceTask>\n`;
        break;
      }
      case 'scriptTask': {
        let script = "";
        if(node.script) {
          script = `${tab(6)}<script><![CDATA[${node.script}]]></script>\n`;
        }
        processXML += `${tab(4)}<scriptTask id="${node.id}" name="${node.label}">\n${script}${tab(4)}</scriptTask>\n`;
        break;
      }
      case 'receiveTask':
        processXML += `${tab(4)}<receiveTask id="${node.id}" name="${node.label}"></receiveTask>\n`;
        break;
      case 'mailTask': {
        let to = `${tab(8)}<flowable:field name="to">\n`;
        to += `${tab(10)}<flowable:string><![CDATA[${node.to}]]></flowable:string>\n`;
        to += `${tab(8)}</flowable:field>\n`;
        let subject = `${tab(8)}<flowable:field name="subject">\n`;
        subject += `${tab(10)}<flowable:string><![CDATA[${node.subject}]]></flowable:string>\n`;
        subject += `${tab(8)}</flowable:field>\n`;
        let text = `${tab(8)}<flowable:field name="text">\n`;
        text += `${tab(10)}<flowable:string><![CDATA[${node.content}]]></flowable:string>\n`;
        text += `${tab(8)}</flowable:field>\n`;
        let extension = `${tab(6)}<extensionElements>\n${to}${subject}${text}${tab(6)}</extensionElements>\n`;
        processXML += `${tab(4)}<serviceTask id="${node.id}" name="${node.label}" flowable:type="mail">\n${extension}${tab(4)}</serviceTask>\n`;
        break;
      }
      case 'timerStart': {
        const timer = `${tab(6)}<timerEventDefinition>\n${tab(8)}<timeCycle>${node.cycle}</timeCycle>\n${tab(6)}</timerEventDefinition>\n`;
        processXML += `${tab(4)}<startEvent id="${node.id}" isInterrupting="false">\n${timer}${tab(4)}</startEvent>\n`;
        break;
      }
      case 'timerCatch': {
        const timer = `${tab(6)}<timerEventDefinition>\n${tab(8)}<timeCycle>${node.cycle}</timeCycle>\n${tab(6)}</timerEventDefinition>\n`;
        processXML += `${tab(4)}<intermediateCatchEvent id="${node.id}">\n${timer}${tab(4)}</intermediateCatchEvent>\n`;
        break;
      }
      case 'signalStart': {
        const signal = `${tab(6)}<signalEventDefinition signalRef="${node.signal}"></signalEventDefinition>\n`;
        processXML += `${tab(4)}<startEvent id="${node.id}" isInterrupting="true">\n${signal}${tab(4)}</startEvent>\n`;
        break;
      }
      case 'signalCatch': {
        const signal = `${tab(6)}<signalEventDefinition signalRef="${node.signal}"></signalEventDefinition>\n`;
        processXML += `${tab(4)}<intermediateCatchEvent id="${node.id}">\n${signal}${tab(4)}</intermediateCatchEvent>\n`;
        break;
      }
      case 'messageStart': {
        const message = `${tab(6)}<messageEventDefinition messageRef="${node.message}"></messageEventDefinition>\n`;
        processXML += `${tab(4)}<startEvent id="${node.id}" isInterrupting="true">\n${message}${tab(4)}</startEvent>\n`;
        break;
      }
      case 'messageCatch': {
        const message = `${tab(6)}<messageEventDefinition messageRef="${node.message}"></messageEventDefinition>\n`;
        processXML += `${tab(4)}<intermediateCatchEvent id="${node.id}">\n${message}${tab(4)}</intermediateCatchEvent>\n`;
        break;
      }
      case 'gateway':
        processXML += `${tab(4)}<exclusiveGateway id="${node.id}" name="${node.label}"></exclusiveGateway>\n`;
        break;
      case 'exclusiveGateway':
        processXML += `${tab(4)}<exclusiveGateway id="${node.id}" name="${node.label}"></exclusiveGateway>\n`;
        break;
      case 'parallelGateway':
        processXML += `${tab(4)}<parallelGateway id="${node.id}" name="${node.label}"></parallelGateway>\n`;
        break;
      case 'inclusiveGateway':
        processXML += `${tab(4)}<inclusiveGateway id="${node.id}" name="${node.label}"></inclusiveGateway>\n`;
        break;
      default:
        break;
    }
  });
  json.edges.forEach(edge => {
    BPMNEdge += `${tab(6)}<bpmndi:BPMNEdge bpmnElement="${edge.source}_${edge.sourceAnchor}-${edge.target}_${edge.targetAnchor}" `+
      `id="BPMNEdge_${edge.source}_${edge.sourceAnchor}-${edge.target}_${edge.targetAnchor}">\n`+
        `${tab(8)}<omgdi:waypoint x="${edge.startPoint.x}" y="${edge.startPoint.y}"></omgdi:waypoint>\n`+
        `${tab(8)}<omgdi:waypoint x="${edge.endPoint.x}" y="${edge.endPoint.y}"></omgdi:waypoint>\n`+
      `${tab(6)}</bpmndi:BPMNEdge>\n`;
    let condition = "";
    if(edge.conditionExpression){
      condition = `${tab(6)}<conditionExpression xsi:type="tFormalExpression"><![CDATA[${edge.conditionExpression}]]></conditionExpression>\n`;
    }
    processXML += `${tab(4)}<sequenceFlow id="${edge.source}_${edge.sourceAnchor}-${edge.target}_${edge.targetAnchor}" sourceRef="${edge.source}" targetRef="${edge.target}">${condition}</sequenceFlow>\n`;
  });
  processXML += `${tab(2)}</process>\n`;

  let BPMNDiagram = `${tab(2)}<bpmndi:BPMNDiagram id="BPMNDiagram_${id}">\n`+
      `${tab(4)}<bpmndi:BPMNPlane bpmnElement="${id}" id="BPMNPlane_${id}">\n${BPMNShape}${BPMNEdge}${tab(4)}</bpmndi:BPMNPlane>\n`+
      `${tab(2)}</bpmndi:BPMNDiagram>\n`;

  let xml = `<?xml version="1.0" encoding="UTF-8"?>\n`;
  xml += `<definitions xmlns="http://www.omg.org/spec/BPMN/20100524/MODEL" `+
    `xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" `+
    `xmlns:xsd="http://www.w3.org/2001/XMLSchema" `+
    `xmlns:flowable="http://flowable.org/bpmn" `+
    `xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" `+
    `xmlns:omgdc="http://www.omg.org/spec/DD/20100524/DC" `+
    `xmlns:omgdi="http://www.omg.org/spec/DD/20100524/DI" `+
    `typeLanguage="http://www.w3.org/2001/XMLSchema" `+
    `expressionLanguage="http://www.w3.org/1999/XPath" `+
    `targetNamespace="http://www.flowable.org/processdef">\n`;
  xml += signals;
  xml += messages;
  xml += processXML;
  xml += BPMNDiagram;
  xml += `</definitions>`;
  if(createFile) {
    downloadFile(xml, 'application/xml;charset=utf-8;', `${name}.bpmn20.xml`)
  }
  return xml;
}

export function exportImg(canvasPanel,filename,createFile = true) {
  filename = filename || 'flow'
  let canvas = canvasPanel.querySelector('canvas')
  let context = canvas.getContext('2d')

  let imgData = context.getImageData(0, 0, canvas.width, canvas.height).data
  let left = canvas.width;
  let right = 0;
  let top = canvas.height;
  let bottom = 0
  for (let i = 0; i < canvas.width; i++) {
    for (let j = 0; j < canvas.height; j++) {
      let pos = (i + canvas.width * j) * 4
      if (imgData[pos] > 0 || imgData[pos + 1] > 0 || imgData[pos + 2] || imgData[pos + 3] > 0) {
        bottom = Math.max(j, bottom) // 找到有色彩的最下端
        right = Math.max(i, right) // 找到有色彩的最右端
        top = Math.min(j, top) // 找到有色彩的最上端
        left = Math.min(i, left) // 找到有色彩的最左端
      }
    }
  }
  let c = document.createElement('canvas')
  // 四周空白余量
  let blankWidth = 60
  c.width = right - left + blankWidth*2
  c.height = bottom - top + blankWidth*2
  let ctx = c.getContext('2d')
  // 设置白底
  ctx.fillStyle = '#fff';
  ctx.fillRect(0, 0, c.width, c.height);
  ctx.drawImage(canvas, left-blankWidth, top-blankWidth, c.width, c.height, 0, 0, c.width, c.height)
  let data = c.toDataURL("image/jpeg")
  if(createFile) {
    let parts = data.split(';base64,');
    let contentType = parts[0].split(':')[1];
    let raw = window.atob(parts[1]);
    let uInt8Array = new Uint8Array(raw.length);
    for (let i = 0; i < raw.length; ++i) {
      uInt8Array[i] = raw.charCodeAt(i);
    }
    downloadFile(uInt8Array, contentType, `${filename}.jpg`)
  }
  return data
}

function downloadFile(data, type, filename) {
  const blob = new Blob([data], { type });
  let link = document.createElement('a');
  if (link.download !== undefined) {
    let url = URL.createObjectURL(blob);
    link.setAttribute('href', url);
    link.setAttribute('download', filename);
    link.style.visibility = 'hidden';
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  }
}