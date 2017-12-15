package org.plugins.qmstr.QuartermasterAction

import static org.plugins.qmstr.QuartermasterAction.DescriptorImpl.QMSTRQ_PROJECT_BLOCK_NAME

def f = namespace(lib.FormTagLib);

f.optionalBlock(name: QMSTRQ_PROJECT_BLOCK_NAME, title: _('quartermaster.project'), checked: instance != null) {
    f.entry(field:'', title: _('qmstr_param')) {
        f.textbox()
    }
}
