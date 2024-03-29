package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type GoalDelegationHooks struct {
	k         Keeper
	missionID uint64
}

// NewGoalDelegationHooks returns a StakingHooks that triggers mission completion on delegation for an account
func (k Keeper) NewGoalDelegationHooks(missionID uint64) GoalDelegationHooks {
	return GoalDelegationHooks{k, missionID}
}

var _ stakingtypes.StakingHooks = GoalDelegationHooks{}

// BeforeDelegationCreated completes mission when a delegation is performed
func (h GoalDelegationHooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, _ sdk.ValAddress) error {
	// TODO handle error
	_, _ = h.k.CompleteGoal(ctx, h.missionID, delAddr.String())
	return nil
}

// AfterValidatorCreated implements StakingHooks
func (h GoalDelegationHooks) AfterValidatorCreated(_ sdk.Context, _ sdk.ValAddress) error {
	return nil
}

// AfterValidatorRemoved implements StakingHooks
func (h GoalDelegationHooks) AfterValidatorRemoved(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

// BeforeDelegationSharesModified implements StakingHooks
func (h GoalDelegationHooks) BeforeDelegationSharesModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

// AfterDelegationModified implements StakingHooks
func (h GoalDelegationHooks) AfterDelegationModified(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}

// BeforeValidatorSlashed implements StakingHooks
func (h GoalDelegationHooks) BeforeValidatorSlashed(_ sdk.Context, _ sdk.ValAddress, _ sdk.Dec) error {
	return nil
}

// BeforeValidatorModified implements StakingHooks
func (h GoalDelegationHooks) BeforeValidatorModified(_ sdk.Context, _ sdk.ValAddress) error {
	return nil
}

// AfterValidatorBonded implements StakingHooks
func (h GoalDelegationHooks) AfterValidatorBonded(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

// AfterValidatorBeginUnbonding implements StakingHooks
func (h GoalDelegationHooks) AfterValidatorBeginUnbonding(_ sdk.Context, _ sdk.ConsAddress, _ sdk.ValAddress) error {
	return nil
}

// BeforeDelegationRemoved implements StakingHooks
func (h GoalDelegationHooks) BeforeDelegationRemoved(_ sdk.Context, _ sdk.AccAddress, _ sdk.ValAddress) error {
	return nil
}
