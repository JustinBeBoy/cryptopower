package localizable

const ENGLISH = "en"

// one string per line, no multiline
// semicolon is not compulsory
const EN = `
"appName" = "cryptopower";
"appTitle" = "cryptopower (%s)";
"recentTransactions" = "Recent Transactions";
"recentProposals" = "Recent Proposals";
"seeAll" = "See all";
"send" = "Send";
"receive" = "Receive";
"online" = "Online, ";
"offline" = "Offline, ";
"viewDetails" = "View details";
"hideDetails" = "Hide details";
"peers" = "peers";
"connectedPeersCount" = "Connected peers count";
"noConnectedPeer" = "no connected peers.";
"disconnect" = "Disconnect";
"reconnect" = "Reconnect";
"currentTotalBalance" = "Current Total Balance";
"totalBalance" = "Total Balance";
"walletStatus" = "Wallet Status:";
"blockHeaderFetched" = "Block header fetched";
"noTransactions" = "No transactions";
"fetchingBlockHeaders" = "Fetching block headers · %v%%";
"discoveringWalletAddress" = "Discovering wallet address · %v%%";
"rescanningHeaders" = "Rescanning headers · %v%%";
"rescanningBlocks" = "Rescanning blocks";
"blocksScanned" = "Blocks scanned";
"blocksLeft" = "%d blocks left"
"syncSteps" = "Step %d/3";
"blockHeaderFetchedCount" = "%d of %d";
"timeLeft" = "%v left";
"connectedTo" = "connected to %v peers";
"connecting" = "Connecting...";
"synced" = "Synced";
"sync" = "Sync";
"autoSyncInfo" = "Auto sync feature has been enable, and wallets are not synced.\nWould you like to start syncing your wallets now?";
"syncingState" = "Syncing...";
"waitingState" = "Waiting...";
"walletNotSynced" = "Not Synced";
"cancel" = "Cancel";
"resumeAccountDiscoveryTitle" = "Unlock to resume restoration";
"unlockWithPassword" = "Unlock with password"
"unlock" = "Unlock";
"syncingProgress" = "Syncing progress";
"syncingProgressStat" = "%s behind";
"noWalletLoaded" = "No wallet loaded";
"lastBlockHeight" = "Last Block Height";
"ago" = "ago";
"newest" = "Newest";
"oldest" = "Oldest";
"all" = "All";
"transferred" = "Transferred";
"sent" = "Sent";
"received" = "Received";
"yourself" = "Yourself";
"mixed" = "Mixed";
"unmined" = "Unmined";
"immature" = "Immature";
"voted" = "Voted";
"revoked" = "Revoked";
"live" = "Live";
"expired" = "Expired";
"purchased" = "Purchased";
"revocation" = "Revocation";
"rebroadcast" = "Rebroadcast";
"staking" = "Staking";
"immatureRewards" = "Immature Rewards";
"lockedByTickets" = "Locked By Tickets";
"immatureStakeGen" = "Immature Stake Gen";
"votingAuthority" = "Voting Authority";
"unknown" = "Unknown";
"nConfirmations" = "%d Confirmations";
"from" = "From";
"to" = "To";
"fee" = "Fee";
"includedInBlock" = "Included in block";
"type" = "Type";
"transactionId" = "Transaction ID";
"transactionDetails" = "Transaction details";
"ticketDetails" = "Ticket details"
"xInputsConsumed" = "%d Inputs consumed";
"xOutputCreated" = "%d Outputs created";
"viewOnExplorer" = "View on block explorer";
"watchOnlyWallets" = "Watch-only wallets";
"signMessage" = "Sign message";
"verifyMessage" = "Verify message";
"message" = "Message";
"viewProperty" = "View property";
"stakeShuffle" = "StakeShuffle";
"rename" = "Rename";
"accRenamed" = "Account renamed"
"renameWalletSheetTitle" = "Change wallet name";
"settings" = "Settings";
"createANewWallet" = "Create a new wallet"
"importExistingWallet" = "Import an existing wallet";
"importWatchingOnlyWallet" = "Import a watch-only wallet";
"create" = "Create";
"watchOnlyWalletImported" = "Watch only wallet imported";
"addNewAccount" = "Add account";
"notBackedUp" = "Backup needed";
"labelSpendable" = "Spendable";
"backupSeedPhrase" = "Back up seed word";
"verifySeedInfo" = "Verify your seed word backup so you can recover your funds when needed.";
"createNewAccount" = "Create new account";
"invalidPassphrase" = "Password entered was not valid.";
"passwordNotMatch" = "Passwords do not match"
"Import" = "Import";
"spendingPasswordInfo" = "A spending password helps secure your wallet transactions."
"spendingPasswordInfo2" = "This spending password is for the new wallet only"
"spendingPassword" = "Spending password";
"enterSpendingPassword" = "Enter spending password"
"confirmSpendingPassword" = "Confirm spending password";
"changeSpendingPass" = "Change spending password";
"currentSpendingPassword" = "Current spending password";
"newSpendingPassword" = "New spending password";
"confirmNewSpendingPassword" = "Confirm new spending password";
"spendingPasswordUpdated" = "Spending password updated";
"notifications" = "Notifications";
"beepForNewBlocks" = "Beep for new blocks";
"debug" = "Debug";
"rescanBlockchain" = "Rescan blockchain";
"dangerZone" = "Danger zone";
"removeWallet" = "Remove wallet from device";
"change" = "Change";
"notConnected" = "Not connected to decred network";
"rescanProgressNotification" = "Check progress in overview.";
"rescan" = "Rescan";
"rescanInfo" = "Rescanning may help resolve some balance errors. This will take some time, as it scans the entire blockchain for transactions"
"confirmToRemove" = "Confirm to remove";
"remove" = "Remove";
"confirm" = "Confirm";
"general" = "General";
"unconfirmedFunds" = "Spend Unconfirmed Funds";
"confirmed" = "Confirmed";
"exchangeRate" = "Fetch exchange rate";
"language" = "Language";
"security" = "Security";
"newStartupPass" = "New startup password"
"confirmNewStartupPass" = "Confirm new startup password"
"startupPassConfirm" = "Startup password changed"
"setupStartupPassword" = "Set up startup password"
"startupPasswordInfo" = "Startup password helps protect your wallet from unauthorized access."
"confirmStartupPass" = "Confirm current startup password"
"currentStartupPass" = "Current startup password"
"startupPassword" = "Startup password";
"changeStartupPassword" = "Change startup password";
"startupPasswordEnabled" = "Startup password %v"
"connection" = "Connection";
"connectToSpecificPeer" = "Connect to specific peer";
"changeSpecificPeer" = "Change specific peer";
"peer" = "Peer"
"CustomUserAgent" = "Custom user agent";
"userAgentSummary" = "For exchange rate fetching";
"changeUserAgent" = "Change user agent";
"createStartupPassword" = "Create a startup password";
"confirmRemoveStartupPass" = "Confirm to turn off startup password";
"userAgentDialogTitle" = "Set up user agent";
"overview" = "Overview";
"transactions" = "Transactions";
"wallets" = "Wallets";
"tickets" = "Tickets";
"ticket" = "Ticket";
"more" = "More";
"english" = "English";
"french" = "French";
"spanish" = "Spanish";
"usdBittrex" = "USD (Bittrex)";
"none" = "None";
"proposals" = "Proposals";
"dex" = "Dex";
"governance" = "Governance";
"pending" = "Pending";
"vote" = "Vote";
"revoke" = "Revoke";
"maturity" = "Maturity";
"yesterday" = "Yesterday";
"days" = "Days";
"hours" = "Hours";
"mins" = "Mins";
"secs" = "Secs";
"weekAgo" = "%d week ago";
"weeksAgo" = "%d weeks ago";
"yearAgo" = "%d year ago";
"yearsAgo" = "%d years ago";  
"monthAgo" = "%d month ago";
"monthsAgo" = "%d months ago";
"dayAgo" = "%d day ago";
"daysAgo" = "%d days ago";
"hourAgo" = "%d hour ago";
"hoursAgo" = " %d hours ago";
"minuteAgo" = "%d minute ago";
"minutesAgo" = "%d minutes ago";
"justNow" = "Just now";
"imawareOfRisk" = "I understand the risks";
"unmixedBalance" = "Unmixed balance";
"backupLater" = "Backup later";
"backupNow" = "Backup now";
"status" = "Status";
"daysToVote" = "Days to vote";
"reward" = "Reward";
"viewTicket" = "View associated ticket";
"external" = "External"
"republished" = "Republished unmined transactions to the decred network";
"copied" = "Copied";
"txHashCopied" = "Transaction Hash copied";
"addressCopied" = "Address copied";
"address" = "Address";
"acctDetailsKey" = "%d external, %d internal, %d imported";
"acctNum" = "Account Number";
"acctName" = "Account name";
"acctRenamed" = "Account renamed";
"acctCreated" = "Account created";
"renameAcct" = "Rename account";
"acctCreated" = "Account created"
"hdPath" = "HD Path";
"key" = "Key";
"validateWalSeed" = "Validate wallet seeds";
"clearAll" = "Clear all";
"restoreWallet" = "Restore wallet";
"restoreExistingWallet" = "Restore existing wallet";
"enterSeedPhrase" = "Enter your seed phrase"
"invalidSeedPhrase" = "Invalid seed phrase"
"seedAlreadyExist" = "A wallet with an identical seed already exists."
"walletExist" = "Wallet with name: %s already exist"
"walletNotExist" = "Wallet with ID: %v does not exist"
"invalidHex"     = "Invalid hex"
"restoreWithHex" = "Restore wallet using hex"
"enterHex"       = "Enter Hex"
"submit"         = "Submit"
"walletRestored" = "Wallet restored"
"enterWalletDetails" = "Enter wallet details"
"copy" = "Copy"
"howToCopy" = "How to copy"
"enterAddressToSign" = "Enter an address and message to sign:"
"signCopied" = "Signature copied"
"signature" = "Signature"
"confirmToSign" = "Confirm to sign"
"enterValidAddress" = "Please enter a valid address"
"enterValidMsg" = "Please enter a valid message to sign"
"invalidAddress" = "Invalid address"
"validAddress" = "Valid address"
"addrNotOwned" = "Address not owned by any wallet"
"delete" = "Delete"
"walletName" = "Wallet name"
"enterWalletName" = "Enter wallet name"
"walletRenamed" = "Wallet renamed"
"walletCreated" = "Wallet created"
"addWallet" = "Add wallet"
"selectWallet" = "Select Wallet"
"checkMixerStatus" = "Check mixer status"
"walletRestoreMsg" = "You can restore this wallet from seed word after it is deleted."
"walletRemoved" = "Wallet removed"
"walletRemoveInfo" = "Make sure to have the seed word backed up before removing the wallet"
"watchOnlyWalletRemoveInfo" = "The watch-only wallet will be removed from your app"
"gotIt" = "Got it"
"noValidAccountFound" = "no valid account found"
"mixer" = "Mixer"
"readyToMix" = "Ready to mix"
"mixerRunning" = "Mixer is running..."
"useMixer" = "How to use the mixer?"
"keepAppOpen" = "Keep this app opened"
"mixerShutdown" = "The mixer will automatically stop when unmixed balance are fully mixed."
"votingPreference" = "Voting Preference:"
"noAgendaYet" = "No agendas yet"
"fetchingAgenda" = "Fetching agendas..."
"updatePreference" = "Update Preference"
"approved" = "Approved"
"voting" = "Votación"
"rejected" = "Rejected"
"abandoned" = "Abandoned"
"inDiscussion" = "In discussion"
"fetchingProposals" = "Fetching proposals..."
"fetchProposals" = "Fetch proposals"
"underReview" = "Under review"
"noProposal" = "No proposals %v"
"waitingForAuthor" = "Waiting for author to authorize voting"
"waitingForAdmin" = "Waiting for admin to trigger the start of voting"
"voteTooltip" = "%d %% Yes votes required for approval"
"yes" = "Yes"
"no" = "No"
"totalVotes" = "Total votes:  %6.0f"
"totalVotesReverse" = "%d Total votes"
"quorumRequirement" = "Quorum requirement:  %6.0f"
"discussions" = "Discussions:   %d comments"
"published" = "Published:   %s"
"token" = "Token:   %s"
"proposalVoteDetails" = "Proposal vote details"
"votingServiceProvider" = "Voting service provider"
"selectVSP" = "Select a VSP..."
"addVSP" = "Add a new VSP..."
"save" = "Save"
"noVSPLoaded" = "No vsp loaded. Check internet connection and try again."
"extendedPubKey" = "Extended public key"
"enterXpubKey" = "Enter a valid extended PubKey"
"xpubKeyErr" = "Error checking xpub: %v"
"xpubWalletExist" = "A wallet with an identical extended public key already exists."
"hint" = "Hint"
"addAcctWarn" = "%v Accounts %v cannot %v be deleted once created.%v"
"verifyMessageInfo" = "%v You can use this form to verify the signature's validity after you or your counterparty have generated one.%v After you've input the address, message, and signature, you'll see VALID if the signature matches the address and message correctly, and INVALID otherwise.%v"
"txDetailsInfo" = "%v Tap on %v blue text %v to copy the item %v"
"setupMixerInfo" = "%v Two dedicated accounts %v mixed %v & %v unmixed %v will be created in order to use the mixer. %v This action cannot be undone.%v"
"backupInfo" = "%v No backup - no coins! %v In order not to lose your coins when your device is lost or broken, please make a wallet backup %v Now %v and keep it in %v a safe palce! %v"
"signMessageInfo" = "%v Signing a message with an address' private key allows you to prove that you are the owner of a given address to a possible counterparty.%v"
"privacyInfo" = "%v When the mixer is activated, funds will be gradually transfered from the unmixed account to the mixed account. %v Important: keep this app open while mixer is running. %v The mixer routine will automatically stop when the unmixed balance is fully mixed.%v"
"allowUnspendUnmixedAcct" = "%v Spendings from unmixed accounts could potentially be traced back to you %v Please type %v I understand the risks %v to allow spending from unmixed accounts.%v"
"balToMaintain" = "Balance to maintain (DCR)"
"autoTicketPurchase" = "Auto ticket purchase"
"purchasingAcct" = "Purchasing account"
"locked" = "Locked"
"balance" = "Balance"
"allTickets" = "All tickets"
"noTickets" = "No tickets yet"
"noActiveTickets" = "No active tickets"
"liveTickets" = "Live Tickets"
"ticketRecord" = "Ticket Record"
"rewardsEarned" = "Rewards Earned"
"noReward" = "Stakey sees no rewards"
"loadingPrice" = "Loading price"
"notAvailable" = "Not available"
"ticketConfirmed" = "Ticket(s) Confirmed"
"backStaking" = "Back to staking"
"ticketSettingSaved" = "Auto ticket purchase setting saved successfully."
"autoTicketWarn" = "Settings can not be modified when ticket buyer is running."
"autoTicketInfo" = "Cryptopower must remain running, for tickets to be automatically purchased"
"confirmPurchase" = "Confirm Automatic Ticket Purchase"
"ticketError" = "Ticket buyer account error: %v"
"walletToPurchaseFrom" = "Wallet to purchase from: %s"
"selectedAcct" = "Selected account: %s"
"balToMaintainValue" = "Balance to maintain: %2.f"
"stake" = "Stake"
"total" = "Total"
"insufficentFund" = "Insufficient funds"
"ticketPrice" = "Ticket Price"
"unminedInfo" = "Broadcasted %v"
"immatureInfo" = "Mature in %v of %v blocks (%v)"
"liveInfo" = "Waiting to be chosen to vote"
"liveInfoDisc" = "live for %v days, (%v of %v remaining)"
"liveInfoDiscSub" = "There is a 0.5% chance of expiring before being chosen to vote (this expiration returns the original Stake price without a reward)"
"votedInfo" = "Congratulations! This Stake has voted."
"votedInfoDisc" = "The Stake price + reward will become spendable after %d blocks (~%s)"
"revokeInfo" = "This Stake has been revoked."
"revokeInfoDisc" = "The Stake price will become spendable after %d blocks (~%s)"
"expiredInfo" = "This Stake has not been chosen to vote within %d blocks, and thus expired"
"expiredInfoDisc" = "Expired tickets will be revoked to return the original Stake price to you"
"expiredInfoDiscSub" = "If a Stake is not revoked automatically, use the revoke button."
"liveIn" = "Live in"
"spendableIn" = "Spendable in"
"duration" = "%s (%d/%d blocks)"
"daysToMiss" = "Days to miss"
"stakeAge" = "Stake age"
"selectOption" = "Select one of the options below to vote"
"updateVotePref" = "Update Voting Preference"
"voteUpdated" = "Vote preference updated successfully"
"votingWallet" = "Voting wallet"
"voteConfirm" = "Confirm to vote"
"voteSent" = "Vote sent successfully, refreshing proposals!"
"numberOfVotes" = "You have %d votes"
"notEnoughVotes" = "You don’t have enough votes"
"search" = "Search"
"consensusChange" = "Consensus Changes"
"onChainVote" = "On-chain voting for upgrading the Decred network consensus rules."
"offChainVote" = "Off-chain voting for development and marketing initiatives funded by the Decred treasury."
"consensusDashboard" = "Consensus Vote Dashboard"
"copyLink" = "Copy and paste the link below in your browser."
"webURL" = "Web URL"
"votingDashboard" = "Voting Dashboard"
"updated" = "Updated"
"viewOnPoliteia" = "View on Politeia"
"votingInProgress" = "Voting in progress..."
"version" = "Version"
"published2" = "Published"
"howGovernanceWork" = "How does Governance Work?"
"governanceInfo" = "%vThe Decred community can participate in proposal discussions for new initiatives and request funding for these initiatives. Decred stakeholders can vote if these proposals should be approved and paid for by the Decred Treasury. %v Would you like to fetch and view the proposals?%v"
"proposalInfo" = "Proposals and politeia notifications can be enabled or disabled from the settings page."
"selectTicket" = "Select a ticket to vote"
"hash" = "Hash"
"max" = "MAX"
"noValidWalletFound" = "no valid wallet found"
"securityTools" = "Security tools"
"about" = "About"
"help" = "Help"
"darkMode" = "Dark mode"
"txNotification" = "Transaction notification %s"
"propNotification" = "Proposal notification %s"
"propNotif" = "Proposal notification"
"httpReq" = "For HTTP request"
"enabled" = "enabled"
"disable" = "Disable"
"disabled" = "disabled"
"governanceSettingsInfo" = "Are you sure you want to disable governance? This will clear all available proposals"
"propFetching" = "Proposals fetching %s. %s"
"checkGovernace" = "Check Governance page"
"removePeer" = "Remove specific peer"
"removePeerWarn" = "Are you sure you want to proceed with removing the specific peer?"
"removeUserAgent" = "Remove user agent"
"removeUserAgentWarn" = "Are you sure you want to proceed with removing the user agent?"
"ipAddress" = "IP address"
"userAgent" = "User agent"
"validateMsg" = "Validate address"
"validate" = "Validate"
"helpInfo" = "For more information, click and please visit the Decred documentation."
"documentation" = "Documentation"
"verifyMsgNote" = "Enter the address, signature, and message to verify:"
"verifyMsgError" = "Error verifying message: %v"
"invalidSignature" = "Invalid signature or message"
"validSignature" = "Valid signature"
"emptySign" = "Field cannot be empty. Please provide valid signature."
"emptyMsg" = "Field cannot be empty. Please provide valid signed message."
"clear" = "Clear"
"validateAddr" = "Validate address"
"validateNote" = "Enter an address to validate:"
"owned" = "Valid address owned by you."
"notOwned" = "Valid address not owned by you."
"buildDate" = "Build date"
"network" = "Network"
"license" = "License"
"checkWalletLog" = "Check wallet logs"
"checkStatistics" = "Check statistics"
"statistics" = "Statistics"
"confirmDexReset" = "Confirm DEX Client Reset"
"dexStartupErr" = "Unable to start DEX client: %v"
"dexResetInfo" = "You may need to restart cryptopower before you can use the DEX again. Proceed?"
"resetDexClient" = "Reset DEX Client"
"walletLog" = "Wallet log"
"build" = "Build"
"peersConnected" = "Peers connected"
"uptime" = "Uptime"
"bestBlocks" = "Best block"
"bestBlockTimestamp" = "Best block timestamp"
"bestBlockAge" = "Best block age"
"walletDirectory" = "Wallet data directory"
"dateSize" = "Wallet data"
"exit" = "Exit"
"loading" = "Loading..."
"openingWallet" = "Opening wallets"
"welcomeNote" = "Welcome to Cryptopower Wallet, a secure & open-source wallet."
"generateAddress" = "Generate new address"
"receivingAddress" = "Receiving account"
"yourAddress" = "Your Address"
"receiveInfo" = "Each time you receive a payment, a new address is generated to protect your privacy."
"dcrReceived" = "You have received %s DCR"
"ticektVoted" =  "A ticket just voted\nVote reward: %s DCR"
"ticketRevoked" = "A ticket was revoked"
"proposalAddedNotif" = "A new proposal has been added Name: %s"
"voteStartedNotif" = "Voting has started for proposal with Token: %s"
"voteEndedNotif" = "Voting has ended for proposal with Token: %s"
"newProposalUpdate" = "New update for proposal with Token: %s"
"walletSyncing" = "Wallet is syncing, please wait"
"next" = "Next"
"retry" = "Retry"
"estimatedTime" = "Estimated time"
"estimatedSize" = "Estimated size"
"rate" = "Rate"
"totalCost" = "Total cost"
"balanceAfter" = "Balance after send"
"sendingAcct" = "Sending account"
"txEstimateErr" = "Error estimating transaction: %v"
"sendInfo" = "Input or scan the destination wallet address and input the amount to send funds."
"amount" = "Amount"
"txSent" = "Transaction sent!"
"confirmSend" = "Confim to send"
"sendingFrom" = "Sending from"
"sendWarning" = "Your DCR will be sent after this step."
"destAddr" = "Destination Address"
"myAcct" = "My account"
"selectWalletToOpen" = "Select the wallet you would like to open."
"securityToolsInfo" = "%v Various tools that help in different aspects of crypto currency security will be located here. %v"
"continue" = "Continue"
"restore" = "Restore"
"newWallet" = "New wallet"
"selectWalletType" = "Select the type of wallet you want to create"
"whatToCallWallet" = "What would you like to call your wallet?"
"existingWalletName" = "What is your wallet existing wallet name?"
"ticketVotedTitle" = "Ticket, Voted"
"ticketRevokedTitle" = "Ticket, Revoked"
"syncCompTime" = "Est. Sync completion time"
"info" = "Info"
"changeAccount" = "Change account"
"mixedAccount" = "Mixed account"
"coordinationServer" = "Coordination server"
"unmixed" = "unmixed"
"allowSpendingFromUnmixedAccount" = "Allow spending from unmixed account"
"walletSettings" = "Wallet settings"
"syncCompTime" = "Est. Sync completion time"
"info" = "Info"
"changeWalletName" = "Change wallet name"
"account" = "Account"
"selectDexServerToOpen" = "Select the Dex server you would like to open."
"addDexServer" = "Add dex server"
"canBuy" = "Can Buy"
"ok" = "OK"
"treasurySpending" = "Treasury Spending"
"treasurySpendingInfo" = "Spending treasury funds now requires stakeholders to vote on the expenditure. You can participate and set a voting policy for treasury spending by a particular Governance Key. The keys can be verified in the dcrd source."
"verifyGovernanceKeys" = "Verify Governance Keys"
"piKey" = "Pi key"
"noPoliciesYet" = "No policies yet"
"fetchingPolicies" = "Fetching policies"
"setTreasuryPolicy" = "Set treasury policy"
"abstain" = "Abstain"
"confirmVote" = "Confirm your vote"
"policySetSuccessfully" = "Your treasury policy has been successfully updated!"
"colon" = ": "
"removeWalletInfo" = "%v Are you sure you want to remove %v %s%v? Enter the name of the wallet below to verify. %v"
"confirmUmixedSpending" = "Confirm to allow spending from unmixed accounts"
"ok" = "OK"
"accountMixer" = "AccountMixer"
"walletLengthError" = "Wallet name must be less than 32 characters"
"validateHostErr" = "%s is not a valid IP or URL address"
"copyBlockLink" = "Copy block explorer link"
"lifeSpan" = "Life Span"
"votedOn" = "Voted on"
"missedOn" = "Missed on"
"missedTickets"="Missed Ticket"
"revokeCause" = "Revocation cause"
"expiredOn" = "Expired on"
"purchasedOn" = "Purchased On"
"confStatus" = "Confirmation Status"
"txFee" = "Transaction Fee"
"vsp" = "VSP"
"vspFee" = "VSP Fee"
"dexDataReset" = "DEX client data reset complete."
"dexDataResetFalse" = "DEX client data reset failed. Check the logs."
"notSameAccoutMixUnmix" = "Cannot use same account for mixed & unmixed"
"walletNameMismatch" = "Wallet name entered doesn't match selected one"
"confirmPending" = "Confirmation Required"
"multipleMixerAccNeeded" = "Set up mixer by creating two needed accounts"
"initiateSetup" = "Initiate Setup"
"takenAccount" = "Account name is taken"
"mixerAccErrorMsg" = "There are existing accounts named mixed or unmixed. Please change the name to something else for now. You can change them back after the setup."
"backAndRename" = "Go back & rename"
"moveToUnmixed" = "Move funds to unmixed account"
"seedValidationFailed" = "Failed to verify. Please go through every wallet seed and try again."
"invalidAmount" = "Invalid amount"
`
