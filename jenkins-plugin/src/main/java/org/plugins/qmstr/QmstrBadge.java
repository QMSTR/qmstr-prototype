package org.plugins.qmstr;

import hudson.Extension;
import hudson.model.BuildBadgeAction;

@Extension
public class QmstrBadge implements BuildBadgeAction{

    @Override
    public String getIconFileName() {
        // TODO Auto-generated method stub
        return null;
    }

    @Override
    public String getDisplayName() {
        // TODO Auto-generated method stub
        return "Message";
    }
    @Override
    public String getUrlName() {
        // TODO Auto-generated method stub
        return null;
    }
}
