import fs from 'fs';
import path from 'path';

// 递归获取目录下所有的 .vue 文件
const getAllVueFiles = (dir, fileList = []) => {
    const files = fs.readdirSync(dir);
    files.forEach(file => {
        const filePath = path.join(dir, file);
        if (fs.statSync(filePath).isDirectory()) {
            getAllVueFiles(filePath, fileList);
        } else if (filePath.endsWith('.vue')) {
            fileList.push(filePath);
        }
    });
    return fileList;
}

// 从 .vue 文件内容中提取组件名称
const extractComponentName = (fileContent) => {
    const regex = /defineOptions\(\s*{\s*name:\s*["']([^"']+)["']/;
    const match = fileContent.match(regex);
    return match ? match[1] : null;
}

// Vite 插件定义
const vueFilePathPlugin = (outputFilePath) => {
    let root;

    return {
        name: 'vue-file-path-plugin',
        configResolved(resolvedConfig) {
            root = resolvedConfig.root;
        },
        buildEnd() {
            const vueFiles = [
                ...getAllVueFiles(path.join(root, 'src/view')),
                ...getAllVueFiles(path.join(root, 'src/plugin'))
            ];
            const pathNameMap = vueFiles.reduce((acc, filePath) => {
                const content = fs.readFileSync(filePath, 'utf-8');
                const componentName = extractComponentName(content);
                if (componentName) {
                    let relativePath ="/" + path.relative(root, filePath).replace(/\\/g, '/');
                    acc[relativePath] = componentName;
                }
                return acc;
            }, {});
            const outputContent = JSON.stringify(pathNameMap, null, 2);
            // 写入指定文件
            fs.writeFileSync(outputFilePath, outputContent);
        }
    };
}

export default vueFilePathPlugin