package com.github.pratyayganguli.transactions.common;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@AllArgsConstructor
@NoArgsConstructor
@Data
public class PlanDetails {
    // these are constants related to the max number of devices being supported...
    private static final int MAX_DEVICES_FREE = 1;
    private static final int MAX_DEVICES_POCKET = 2;
    private static final int MAX_DEVICES_PREMIUM = 5;

    private PlanType planType;
    private float currency;
    private int currentDevicesLoggedIn;
}
