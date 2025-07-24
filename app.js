// Application state
let currentScore = 0;
let challengeScore = 0;
let completedExercises = new Set();
let roadmapProgress = {};

// Enhanced exercise data
const exercises = [
    {
        question: "Which of the following is a valid way to declare a variable in Go?",
        options: ["Option A only", "Option B only", "Option C only", "All of the above"],
        correct: 3,
        explanation: "All three ways are valid in Go: var with explicit type, short declaration :=, and var with type inference.",
        difficulty: "beginner"
    },
    {
        question: "What is the output of this code?",
        options: ["42: answer", "answer: 42", "Compilation error", "42 answer"],
        correct: 0,
        explanation: "The function returns 42 as the first value and 'answer' as the second. Printf formats them as '42: answer'.",
        difficulty: "beginner"
    },
    {
        question: "What will be the length of the slice after these operations?",
        options: ["2", "3", "4", "5"],
        correct: 1,
        explanation: "slice starts with [1,2,3], becomes [1,2,3,4,5] after append, then [2,3,4] after slice[1:4], so length is 3.",
        difficulty: "beginner"
    },
    {
        question: "Which statement about Go interfaces is correct?",
        options: [
            "MyWriter must explicitly declare it implements Writer",
            "MyWriter automatically implements Writer",
            "MyWriter cannot implement Writer without inheritance",
            "This code will not compile"
        ],
        correct: 1,
        explanation: "Go uses implicit interface satisfaction. Any type that implements all methods of an interface automatically satisfies that interface.",
        difficulty: "intermediate"
    },
    {
        question: "What happens when this code runs?",
        options: ["Prints 1 and exits", "Prints 1, 2, 3", "Deadlock error", "Compilation error"],
        correct: 2,
        explanation: "The channel has buffer size 2, so the first two sends succeed, but the third send blocks because the buffer is full, causing a deadlock.",
        difficulty: "intermediate"
    },
    {
        question: "What is the best practice for this function?",
        options: [
            "Add defer file.Close()",
            "Use panic instead of returning error",
            "Ignore the error",
            "The code is perfect as-is"
        ],
        correct: 0,
        explanation: "Always close opened files. defer file.Close() ensures the file is closed even if an error occurs later.",
        difficulty: "intermediate"
    },
    {
        question: "What is the purpose of context.Context in Go?",
        options: [
            "Only for HTTP requests",
            "Cancellation, deadlines, and request-scoped values",
            "Only for database connections",
            "Only for logging"
        ],
        correct: 1,
        explanation: "Context provides cancellation signals, deadlines, and request-scoped values across API boundaries and goroutines.",
        difficulty: "advanced"
    },
    {
        question: "Which statement about Go's memory management is correct?",
        options: [
            "Go has manual memory management like C",
            "Go uses reference counting for garbage collection",
            "Go uses a concurrent, tri-color mark-and-sweep GC",
            "Go never frees memory automatically"
        ],
        correct: 2,
        explanation: "Go uses a concurrent, tri-color mark-and-sweep garbage collector that runs concurrently with the program.",
        difficulty: "advanced"
    }
];

// Wait for DOM to be fully loaded
document.addEventListener('DOMContentLoaded', function() {
    console.log('DOM loaded, initializing app...');
    initializeApp();
});

function initializeApp() {
    console.log('Initializing app...');
    setupTabNavigation();
    setupOSTabNavigation();
    setupCodeExamples();
    setupCopyButtons();
    setupRoadmapProgress();
    setupExercises();
    loadProgress();
    console.log('App initialized successfully');
}

// Tab Navigation - Fixed
function setupTabNavigation() {
    console.log('Setting up tab navigation...');
    const tabBtns = document.querySelectorAll('.tab-btn');
    const tabContents = document.querySelectorAll('.tab-content');
    
    console.log('Found tab buttons:', tabBtns.length);
    console.log('Found tab contents:', tabContents.length);

    tabBtns.forEach((btn, index) => {
        console.log(`Setting up tab button ${index}:`, btn.dataset.tab);
        btn.addEventListener('click', function(e) {
            e.preventDefault();
            const targetTab = this.dataset.tab;
            console.log('Tab clicked:', targetTab);
            
            // Remove active class from all tabs and contents
            tabBtns.forEach(b => b.classList.remove('active'));
            tabContents.forEach(c => c.classList.remove('active'));
            
            // Add active class to clicked tab and corresponding content
            this.classList.add('active');
            const targetContent = document.getElementById(targetTab);
            if (targetContent) {
                targetContent.classList.add('active');
                console.log('Successfully switched to tab:', targetTab);
            } else {
                console.error('Target content not found:', targetTab);
            }
            
            // Save current tab to localStorage
            try {
                localStorage.setItem('currentTab', targetTab);
            } catch (e) {
                console.warn('Could not save to localStorage:', e);
            }
        });
    });
}

// OS Tab Navigation - Fixed
function setupOSTabNavigation() {
    console.log('Setting up OS tab navigation...');
    const osTabs = document.querySelectorAll('.os-tab');
    const osInstructions = document.querySelectorAll('.os-instructions');
    
    console.log('Found OS tabs:', osTabs.length);
    console.log('Found OS instructions:', osInstructions.length);

    osTabs.forEach(tab => {
        tab.addEventListener('click', function(e) {
            e.preventDefault();
            const targetOS = this.dataset.os;
            console.log('OS tab clicked:', targetOS);
            
            // Remove active class from all OS tabs and instructions
            osTabs.forEach(t => t.classList.remove('active'));
            osInstructions.forEach(i => i.classList.remove('active'));
            
            // Add active class to clicked tab and corresponding instructions
            this.classList.add('active');
            const targetInstructions = document.getElementById(targetOS);
            if (targetInstructions) {
                targetInstructions.classList.add('active');
                console.log('Successfully switched to OS:', targetOS);
            } else {
                console.error('Target OS instructions not found:', targetOS);
            }
        });
    });
}

// Code Examples Expandable Sections - Fixed
function setupCodeExamples() {
    console.log('Setting up code examples...');
    const exampleHeaders = document.querySelectorAll('.example-header');
    
    console.log('Found example headers:', exampleHeaders.length);
    
    exampleHeaders.forEach(header => {
        header.addEventListener('click', function(e) {
            e.preventDefault();
            const targetId = this.dataset.toggle;
            const content = document.getElementById(targetId);
            const icon = this.querySelector('.toggle-icon');
            
            console.log('Example header clicked:', targetId);
            
            if (content && icon) {
                // Toggle expanded state
                this.classList.toggle('expanded');
                content.classList.toggle('expanded');
                
                // Update icon
                if (this.classList.contains('expanded')) {
                    icon.textContent = '▲';
                } else {
                    icon.textContent = '▼';
                }
                
                console.log('Toggled example:', targetId, 'expanded:', this.classList.contains('expanded'));
            } else {
                console.error('Content or icon not found for:', targetId);
            }
        });
    });
}

// Copy to Clipboard Functionality - Fixed
function setupCopyButtons() {
    console.log('Setting up copy buttons...');
    const copyBtns = document.querySelectorAll('.copy-btn');
    
    console.log('Found copy buttons:', copyBtns.length);
    
    copyBtns.forEach(btn => {
        btn.addEventListener('click', function(e) {
            e.preventDefault();
            const targetId = this.dataset.copy;
            const codeElement = document.getElementById(targetId);
            
            console.log('Copy button clicked:', targetId);
            
            if (codeElement) {
                const codeText = codeElement.textContent;
                
                // Copy to clipboard
                if (navigator.clipboard && navigator.clipboard.writeText) {
                    navigator.clipboard.writeText(codeText).then(() => {
                        console.log('Code copied successfully');
                        showCopyNotification('Code copied to clipboard!');
                        
                        // Update button text temporarily
                        const originalText = this.textContent;
                        this.textContent = 'Copied!';
                        setTimeout(() => {
                            this.textContent = originalText;
                        }, 1000);
                    }).catch(err => {
                        console.error('Failed to copy: ', err);
                        showCopyNotification('Failed to copy code', 'error');
                    });
                } else {
                    // Fallback for older browsers
                    try {
                        const textarea = document.createElement('textarea');
                        textarea.value = codeText;
                        document.body.appendChild(textarea);
                        textarea.select();
                        document.execCommand('copy');
                        document.body.removeChild(textarea);
                        showCopyNotification('Code copied to clipboard!');
                    } catch (err) {
                        console.error('Fallback copy failed:', err);
                        showCopyNotification('Copy not supported', 'error');
                    }
                }
            } else {
                console.error('Code element not found:', targetId);
            }
        });
    });
}

// Show copy notification - Fixed
function showCopyNotification(message, type = 'success') {
    // Remove existing notification
    const existingNotification = document.querySelector('.copy-notification');
    if (existingNotification) {
        existingNotification.remove();
    }
    
    // Create new notification
    const notification = document.createElement('div');
    notification.className = `copy-notification ${type}`;
    notification.textContent = message;
    document.body.appendChild(notification);
    
    // Show notification
    setTimeout(() => {
        notification.classList.add('show');
    }, 10);
    
    // Hide notification after 3 seconds
    setTimeout(() => {
        notification.classList.remove('show');
        setTimeout(() => {
            if (notification.parentNode) {
                notification.remove();
            }
        }, 300);
    }, 3000);
}

// Roadmap Progress Tracking - Fixed
function setupRoadmapProgress() {
    console.log('Setting up roadmap progress...');
    const markCompleteButtons = document.querySelectorAll('.mark-complete');
    
    console.log('Found mark complete buttons:', markCompleteButtons.length);
    
    markCompleteButtons.forEach(btn => {
        btn.addEventListener('click', function(e) {
            e.preventDefault();
            const topicId = this.dataset.topic;
            const topicCard = this.closest('.topic-card');
            const progressFill = topicCard.querySelector('.progress-fill');
            
            console.log('Mark complete clicked:', topicId);
            
            if (topicCard.classList.contains('completed')) {
                // Mark as incomplete
                topicCard.classList.remove('completed');
                progressFill.style.width = '0%';
                this.textContent = 'Mark Complete';
                delete roadmapProgress[topicId];
                console.log('Marked as incomplete:', topicId);
            } else {
                // Mark as complete
                topicCard.classList.add('completed');
                progressFill.style.width = '100%';
                this.textContent = 'Completed ✓';
                roadmapProgress[topicId] = true;
                console.log('Marked as complete:', topicId);
            }
            
            saveProgress();
        });
    });
}

function setupExercises() {
    console.log('Setting up exercises...');
    const optionBtns = document.querySelectorAll('.option-btn');
    
    console.log('Found option buttons:', optionBtns.length);
    
    optionBtns.forEach(btn => {
        btn.addEventListener('click', function(e) {
            e.preventDefault();
            const exerciseIndex = parseInt(this.dataset.exercise);
            const selectedOption = parseInt(this.dataset.option);
            const exercise = exercises[exerciseIndex];
            
            console.log('Exercise option clicked:', exerciseIndex, selectedOption);
            
            // Disable all options for this exercise
            const exerciseOptions = document.querySelectorAll(`[data-exercise="${exerciseIndex}"]`);
            exerciseOptions.forEach(option => {
                option.disabled = true;
                option.style.pointerEvents = 'none';
            });
            
            // Show feedback
            const feedbackElement = document.getElementById(`feedback-${exerciseIndex}`);
            
            if (selectedOption === exercise.correct) {
                // Correct answer
                this.classList.add('correct');
                feedbackElement.innerHTML = `
                    <div class="feedback feedback--correct">
                        <strong>✅ Correct!</strong><br>
                        ${exercise.explanation}
                    </div>
                `;
                
                // Update score if not already completed
                if (!completedExercises.has(exerciseIndex)) {
                    currentScore++;
                    completedExercises.add(exerciseIndex);
                    updateScoreDisplay();
                    console.log('Exercise completed correctly, score:', currentScore);
                }
            } else {
                // Incorrect answer
                this.classList.add('incorrect');
                feedbackElement.innerHTML = `
                    <div class="feedback feedback--incorrect">
                        <strong>❌ Incorrect.</strong><br>
                        The correct answer is: <strong>${exercise.options[exercise.correct]}</strong><br>
                        ${exercise.explanation}
                    </div>
                `;
                
                // Highlight correct answer
                exerciseOptions[exercise.correct].classList.add('correct');
                console.log('Exercise answered incorrectly');
            }
            
            saveProgress();
        });
    });

    // Setup difficulty filter
    const filterBtns = document.querySelectorAll('.filter-btn');
    filterBtns.forEach(btn => {
        btn.addEventListener('click', function() {
            filterBtns.forEach(b => b.classList.remove('active'));
            this.classList.add('active');
            
            const difficulty = this.dataset.difficulty;
            filterExercises(difficulty);
        });
    });

    // Setup coding challenges
    setupCodingChallenges();
}

function filterExercises(difficulty) {
    const exerciseCards = document.querySelectorAll('.exercise-card, .challenge-card');
    
    exerciseCards.forEach(card => {
        if (difficulty === 'all' || card.dataset.difficulty === difficulty) {
            card.style.display = 'block';
        } else {
            card.style.display = 'none';
        }
    });
}

function setupCodingChallenges() {
    const runBtns = document.querySelectorAll('.run-code-btn');
    
    runBtns.forEach(btn => {
        btn.addEventListener('click', function() {
            const challengeCard = this.closest('.challenge-card');
            const codeInput = challengeCard.querySelector('.code-input');
            const output = challengeCard.querySelector('.code-output');
            
            // Simulate code execution (in a real app, this would send to a backend)
            const code = codeInput.value;
            output.innerHTML = `
                <div class="code-result">
                    <strong>Code executed:</strong><br>
                    <pre>${code}</pre>
                    <div class="execution-note">
                        💡 In a real environment, this would compile and run your Go code.
                        For now, check your logic and compare with the expected output.
                    </div>
                </div>
            `;
            
            // Mark challenge as attempted
            const challengeIndex = Array.from(document.querySelectorAll('.challenge-card')).indexOf(challengeCard);
            if (code.trim().length > 50) { // Basic check for substantial code
                challengeScore = Math.max(challengeScore, challengeIndex + 1);
                updateScoreDisplay();
                saveProgress();
            }
        });
    });
}

// Update score display
function updateScoreDisplay() {
    const scoreElement = document.getElementById('score');
    if (scoreElement) {
        scoreElement.textContent = currentScore;
    }

    const mcScoreElement = document.getElementById('mc-score');
    const challengeScoreElement = document.getElementById('challenge-score');
    const overallProgress = document.getElementById('overall-progress');
    
    if (mcScoreElement) mcScoreElement.textContent = currentScore;
    if (challengeScoreElement) challengeScoreElement.textContent = challengeScore;
    
    if (overallProgress) {
        const totalQuestions = exercises.length + 3; // 3 coding challenges
        const totalCompleted = currentScore + challengeScore;
        const percentage = (totalCompleted / totalQuestions) * 100;
        overallProgress.style.width = `${percentage}%`;
    }
}

// Reset exercises
function resetExercises() {
    resetAllExercises();
}

function resetAllExercises() {
    console.log('Resetting all exercises...');
    currentScore = 0;
    challengeScore = 0;
    completedExercises.clear();
    
    // Reset all option buttons
    const optionBtns = document.querySelectorAll('.option-btn');
    optionBtns.forEach(btn => {
        btn.classList.remove('correct', 'incorrect');
        btn.disabled = false;
        btn.style.pointerEvents = 'auto';
    });
    
    // Clear all feedback
    const feedbackElements = document.querySelectorAll('.exercise-feedback');
    feedbackElements.forEach(feedback => {
        feedback.innerHTML = '';
    });

    // Reset coding challenges
    const codeInputs = document.querySelectorAll('.code-input');
    const codeOutputs = document.querySelectorAll('.code-output');
    
    codeInputs.forEach(input => {
        input.value = input.placeholder;
    });
    
    codeOutputs.forEach(output => {
        output.innerHTML = '';
    });
    
    updateScoreDisplay();
    saveProgress();
}

// Save progress to localStorage
function saveProgress() {
    const progressData = {
        currentScore: currentScore,
        challengeScore: challengeScore,
        completedExercises: Array.from(completedExercises),
        roadmapProgress: roadmapProgress
    };
    
    try {
        localStorage.setItem('goLearningProgress', JSON.stringify(progressData));
        console.log('Progress saved');
    } catch (e) {
        console.warn('Could not save progress:', e);
    }
}

// Load progress from localStorage
function loadProgress() {
    try {
        const saved = localStorage.getItem('goLearningProgress');
        if (saved) {
            const progressData = JSON.parse(saved);
            
            currentScore = progressData.currentScore || 0;
            challengeScore = progressData.challengeScore || 0;
            completedExercises = new Set(progressData.completedExercises || []);
            roadmapProgress = progressData.roadmapProgress || {};
            
            updateScoreDisplay();
            restoreRoadmapProgress();
            restoreExerciseProgress();
            console.log('Progress loaded');
        }
        
        // Restore current tab
        const currentTab = localStorage.getItem('currentTab');
        if (currentTab) {
            const tabBtn = document.querySelector(`[data-tab="${currentTab}"]`);
            const tabContent = document.getElementById(currentTab);
            
            if (tabBtn && tabContent) {
                document.querySelectorAll('.tab-btn').forEach(btn => btn.classList.remove('active'));
                document.querySelectorAll('.tab-content').forEach(content => content.classList.remove('active'));
                
                tabBtn.classList.add('active');
                tabContent.classList.add('active');
                console.log('Restored tab:', currentTab);
            }
        }
    } catch (e) {
        console.warn('Could not load progress:', e);
    }
}

// Restore roadmap progress
function restoreRoadmapProgress() {
    Object.keys(roadmapProgress).forEach(topicId => {
        const topicCard = document.querySelector(`[data-topic="${topicId}"]`);
        if (topicCard) {
            const progressFill = topicCard.querySelector('.progress-fill');
            const markCompleteBtn = topicCard.querySelector('.mark-complete');
            
            topicCard.classList.add('completed');
            progressFill.style.width = '100%';
            markCompleteBtn.textContent = 'Completed ✓';
        }
    });
}

// Restore exercise progress
function restoreExerciseProgress() {
    completedExercises.forEach(exerciseIndex => {
        const exercise = exercises[exerciseIndex];
        const exerciseOptions = document.querySelectorAll(`[data-exercise="${exerciseIndex}"]`);
        const feedbackElement = document.getElementById(`feedback-${exerciseIndex}`);
        
        // Disable all options
        exerciseOptions.forEach(option => {
            option.disabled = true;
            option.style.pointerEvents = 'none';
        });
        
        // Show correct answer
        exerciseOptions[exercise.correct].classList.add('correct');
        
        // Show feedback
        feedbackElement.textContent = `Correct! ${exercise.explanation}`;
        feedbackElement.className = 'exercise-feedback correct';
    });
}

// Keyboard shortcuts
document.addEventListener('keydown', function(e) {
    // Ctrl/Cmd + number keys for tab navigation
    if ((e.ctrlKey || e.metaKey) && e.key >= '1' && e.key <= '5') {
        e.preventDefault();
        const tabIndex = parseInt(e.key) - 1;
        const tabBtns = document.querySelectorAll('.tab-btn');
        
        if (tabBtns[tabIndex]) {
            tabBtns[tabIndex].click();
        }
    }
    
    // Escape key to close expanded examples
    if (e.key === 'Escape') {
        const expandedHeaders = document.querySelectorAll('.example-header.expanded');
        expandedHeaders.forEach(header => {
            header.click();
        });
    }
});

// Export functions for global access
window.resetExercises = resetExercises;
window.resetAllExercises = resetAllExercises;
window.goLearningApp = {
    resetExercises,
    resetAllExercises,
    saveProgress,
    loadProgress,
    currentScore,
    challengeScore,
    completedExercises,
    roadmapProgress
};

// Error handling
window.addEventListener('error', function(e) {
    console.error('Application error:', e.error);
    showCopyNotification('An error occurred. Please refresh the page.', 'error');
});

// Cleanup on page unload
window.addEventListener('beforeunload', function() {
    saveProgress();
});

console.log('Script loaded successfully');