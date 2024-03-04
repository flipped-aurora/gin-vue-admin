const ticketPanel = {
	"panels": [
		{
			"width": 80,
			"height": 80,
			"paperFooter": 0,
			"paperHeader": 0,
			"paperNumberDisabled": true,
      "printElements": [
				{
					"options": {
						"left": 0,
						"top": 0,
						"height": 10,
						"width": 120,
            "title":"影厅",
						"testData": "3号厅",
						"field": "hall",
						"fontFamily": "sans-serif",
						"color": "#000000",
						"fontSize": 12,
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 80,
						"top": 0,
						"height": 10,
						"width": 130,
						"testData": "2023-12-28 19:16",
            "title":"时间",
						"field": "today",
						"dataType": "datetime",
						"format": "yyyy-MM-dd HH:mm",
						"color": "#000000",
						"fontSize": 12,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 0,
						"top": 15,
						"height": 10,
						"width": 160,
						"testData": "海王2失落的王国",
            "title":"片名",
						"field": "filmName",
						"color": "#000000",
						"fontSize": 12,
						"fontFamily": "sans-serif",
						"fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 190,
						"top": 15,
						"height": 10,
						"width": 120,
						"testData": "3号厅",
						"field": "hall",
						"fontFamily": "sans-serif",
						"color": "#000000",
						"fontSize": 12,
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 0,
						"top": 30,
						"height": 10,
						"width": 93,
            "title":"座号",
						"testData": "2排10座",
						"field": "seat",
						"color": "#000000",
						"fontSize": 12,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 80,
						"top": 30,
						"height": 15,
						"width": 80,
            "title":"票价",
						"field": "price",
						"testData": "35",
						"color": "#000000",
						"fontSize": 12,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 190,
						"top": 30,
						"height": 10,
						"width": 45,
						"testData": "2023-12-28",
						"field": "today",
            "dataType": "datetime",
						"format": "yyyy-MM-dd",
						"color": "#000000",
						"fontSize": 12,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 0,
						"top": 45,
						"height": 10,
						"width": 93,
            "title":"票类",
						"testData": "成人票",
						"field": "type",
						"color": "#000000",
						"fontSize": 12,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 80,
						"top": 45,
						"height": 10,
						"title": "票外费用",
            "field": "sellerFee",
						"width": 90,
 						"fontSize": 12,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 0,
						"top": 60,
						"height": 10,
						"width": 115,
						"testData": "2023-12-28 19:16",
            "title":"售票时间",
						"field": "datetime",
						"dataType": "datetime",
						"format": "yyyy-MM-dd HH:mm",
						"color": "#000000",
						"fontSize": 12,
            "lineHeight": "15",
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 134,
						"top": 60,
						"height": 10,
						"title": "售票员:hyxc",
						"width": 70,
						"color": "#000000",
						"fontSize": 12,
            "lineHeight": "15",
						"fontFamily": "sans-serif",
						"fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 190,
						"top": 90,
						"height": 10,
						"width": 93,
						"testData": "2排10座",
						"field": "seat",
						"color": "#000000",
						"fontSize": 12,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 0,
						"top": 105,
						"height": 85,
						"width": 85,
						"field": "qrcodeUrl",
						"testData": "1fghjkjhdfa1fghjkjhdfafasfadsfasdfasdfasfadsfsadddddsfdsasfddddsfdsasffd231312",
						"testData": "dd",
						"textType": "qrcode"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 95,
						"top": 105,
						"height": 10,
						"width": 80,
						"title": "影票销售方",
            "field": "seller",
            "data": "黄岩星驰国际影城",
            "fontFamily": "sans-serif",
            "fontSize": 10,
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 95,
						"top": 130,
						"height": 10,
						"width": 85,
            "fontSize": 10,
						"title": "影厅服务费",
            "fontFamily": "sans-serif",
            "field": "charge",
            "data": "0.00",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 190,
						"top": 135,
						"height": 10,
						"width": 40,
						"field": "orderNo",
						"fontSize": 10,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 0,
						"top": 195,
						"height": 10,
						"width": 135,
            "title":"票号",
						"field": "orderNo",
						"fontSize": 11,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
        {
					"options": {
						"left": 190,
						"top": 195,
						"height": 10,
						"width": 50,
						"field": "price",
						"color": "#000000",
						"fontSize": 12,
						"fontFamily": "sans-serif",
            "fontWeight": "bolder"
					},
					"printElementType": {
						"type": "text"
					}
				},
      ]
		}
	]
}

export default ticketPanel;